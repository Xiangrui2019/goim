package connect

import (
	"encoding/binary"
	"errors"
	"goim/public/logger"
	"net"
	"time"
)

const (
	TypeLen       = 2                       // 消息类型字节数组长度
	LenLen        = 2                       // 消息长度字节数组长度
	HeadLen       = TypeLen + LenLen        // 消息头部字节数组长度（消息类型字节数组长度+消息长度字节数组长度）
	ContentMaxLen = 4092                    // 消息体最大长度
	BufLen        = ContentMaxLen + HeadLen // 缓冲buffer字节数组长度

)

var (
	ErrOutOfSize       = errors.New("package content out of size") // package的content字节数组过大
	ErrIllegalValueLen = errors.New("illegal package length")      // 违法的包长度
)

// Codec 编解码器，用来处理tcp的拆包粘包
type Codec struct {
	Conn     net.Conn
	ReadBuf  buffer // 读缓冲
	WriteBuf []byte // 写缓冲
}

// newCodec 创建一个编解码器
func NewCodec(conn net.Conn) *Codec {
	return &Codec{
		Conn:     conn,
		ReadBuf:  newBuffer(conn, BufLen),
		WriteBuf: make([]byte, BufLen),
	}
}

// Read 从conn里面读取数据，当conn发生阻塞，这个方法也会阻塞
func (c *Codec) Read() (int, error) {
	return c.ReadBuf.readFromReader()
}

// Decode 解码数据
// Package 代表一个解码包
// bool 标识是否还有可读数据
func (c *Codec) Decode() (*Package, bool, error) {
	var err error
	// 读取数据类型
	typeBuf, err := c.ReadBuf.seek(0, TypeLen)
	if err != nil {
		return nil, false, nil
	}

	// 读取数据长度
	lenBuf, err := c.ReadBuf.seek(TypeLen, HeadLen)
	if err != nil {
		return nil, false, nil
	}

	// 读取数据内容
	valueType := int(binary.BigEndian.Uint16(typeBuf))
	valueLen := int(binary.BigEndian.Uint16(lenBuf))

	// 数据的字节数组长度大于buffer的长度，返回错误
	if valueLen > ContentMaxLen {
		logger.Logger.Error(ErrIllegalValueLen.Error())
		return nil, false, ErrIllegalValueLen
	}

	valueBuf, err := c.ReadBuf.read(HeadLen, valueLen)
	if err != nil {
		return nil, false, nil
	}
	message := Package{Code: valueType, Content: valueBuf}
	return &message, true, nil
}

// Eecode 编码数据,todo 在业务服务器检查消息投递包的大小
func (c *Codec) Eecode(pack Package, duration time.Duration) error {
	contentLen := len(pack.Content)
	if contentLen > ContentMaxLen {
		return ErrOutOfSize
	}

	binary.BigEndian.PutUint16(c.WriteBuf[0:TypeLen], uint16(pack.Code))
	binary.BigEndian.PutUint16(c.WriteBuf[LenLen:HeadLen], uint16(len(pack.Content)))
	copy(c.WriteBuf[HeadLen:], pack.Content[:contentLen])

	c.Conn.SetWriteDeadline(time.Now().Add(duration))
	_, err := c.Conn.Write(c.WriteBuf[:HeadLen+contentLen])
	if err != nil {
		return err
	}
	return nil
}
