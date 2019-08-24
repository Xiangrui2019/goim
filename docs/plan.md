### 增加功能
1.多业务复用  
2.大群读扩散 （暂时取消） 
3.性能提升  
4.消息存储改为mongodb
5.消息编解码放到logic层  
6.复杂消息定义
7.错误类型整理
8.消息的分布式链路追踪，到客户端
### 1.2.0开发计划
1.协议梳理  
2.数据库梳理    
3.代码开发

pb编译命令
protoc --go_out=. public/proto/protocol.proto
 protoc --go_out ../pb/ message.proto
 
logic 调用 connect层，使用rpc

1.服务发现逻辑
上线，通知到所有用户
2. 