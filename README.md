# go-grpc-examples

gRPC 是一个高性能、开源和通用的 RPC 框架，面向移动和 HTTP/2 设计，带来诸如双向流、流控、头部压缩、单 TCP 连接上的多复用请求等特。这些特性使得其在移动设备上表现更好，更省电和节省空间占用。

在 gRPC 里客户端应用可以像调用本地对象一样直接调用另一台不同的机器上服务端应用的方法，使得您能够更容易地创建分布式应用和服务。

gRPC 默认使用 protocol buffers，这是 Google 开源的一套成熟的结构数据序列化机制，它的作用与 XML、json 类似，但它是二进制格式，性能好、效率高（缺点：可读性差）。

## 环境配置

### 1.安装 protobuf

下载地址： https://github.com/protocolbuffers/protobuf/releases

检查安装：

```shell
$ protoc --version                                                                                                                                                                 
libprotoc 3.20.0
```

### 2.安装相关包

安装 golang 的proto编译支持
```shell
go get -u google.golang.org/protobuf
```

安装 gRPC 包
```shell
go get -u google.golang.org/grpc
```

## gRPC4种模式

gRPC主要有4种请求和响应模式，分别是简单模式(Simple RPC)、服务端流式（Server-side streaming RPC）、客户端流式（Client-side streaming RPC）、和双向流式（Bidirectional streaming RPC）。

1. 简单模式(Simple RPC)：客户端发起请求并等待服务端响应。

2. 服务端流式（Server-side streaming RPC）：客户端发送请求到服务器，拿到一个流去读取返回的消息序列。 客户端读取返回的流，直到里面没有任何消息。

3. 客户端流式（Client-side streaming RPC）：与服务端数据流模式相反，这次是客户端源源不断的向服务端发送数据流，而在发送结束后，由服务端返回一个响应。

4. 双向流式（Bidirectional streaming RPC）：双方使用读写流去发送一个消息序列，两个流独立操作，双方可以同时发送和同时接收。

## package 以及 go_package

package用于proto,在引用时起作用;

option go_package用于生成的.pb.go文件,在引用时和生成go包名时起作用

## 其他

1. mustEmbedUnimplementedHelloServiceServer

# 参考

[Bingjian-Zhu / go-grpc-example](https://github.com/Bingjian-Zhu/go-grpc-example)

[grpc tutorial](https://grpc.io/docs/languages/go/basics/)

[指月小筑 grpc](https://www.lixueduan.com/categories/grpc/)