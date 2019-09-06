package connect

import (
	"goim/conf"
	"goim/public/logger"
	"goim/public/transfer"
	"log"
	"net"
	"net/rpc"
)

type ConnectRPCServer struct{}

// DeliverMessage 投点消息
func (s *ConnectRPCServer) DeliverMessage(req transfer.DeliverMessageReq, resp *transfer.DeliverMessageResp) error {
	// 获取设备对应的TCP连接
	ctx := load(req.DeviceId)
	if ctx == nil {
		logger.Sugar.Error("ctx id nil")
		return nil
	}

	// 发送消息
	err := ctx.Codec.Eecode(Package{Code: CodeMessage, Content: req.Bytes}, WriteDeadline)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func StartRPCServer() {
	rpc.Register(new(ConnectRPCServer))
	tcpAddr, err := net.ResolveTCPAddr("tcp", conf.ConnectRPCClientIP)
	if err != nil {
		log.Println(err)
		return
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
