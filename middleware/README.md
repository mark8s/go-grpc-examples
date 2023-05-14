# go-grpc-middleware

上篇介绍了gRPC中TLS认证和自定义方法认证，最后还简单介绍了gRPC拦截器的使用。gRPC自身只能设置一个拦截器，所有逻辑都写一起会比较乱。本篇简单介绍go-grpc-middleware的使用，包括grpc_zap、grpc_auth和grpc_recovery。

go-grpc-middleware封装了认证（auth）, 日志（ logging）, 消息（message）, 验证（validation）, 重试（retries） 和监控（retries）等拦截器。

安装：

```shell
go get github.com/grpc-ecosystem/go-grpc-middleware
```

