# gRPC

- 是Google开源的RPC框架
- 基于HTTP/2标准设计
- 通过protobuf来定义接口，数据被序列化成二进制编码传输，提高效率
- 有四种调用方式：一元调用、服务端/客户端流式调用、双向流式调用

### protobuf 和 json的区别

- protobuf的编码解码比json快
- protobuf的内存占用更少