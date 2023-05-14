# grpc-validator

本篇将介绍grpc_validator，它可以对gRPC数据的输入和输出进行验证。

这里使用第三方插件protoc-gen-validate自动生成验证规则。

```shell
go get github.com/envoyproxy/protoc-gen-validate
```

## 安装和使用

自行下载：
```shell
https://github.com/bufbuild/protoc-gen-validate/releases
```

