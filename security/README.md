# grpc认证
gRPC建立在HTTP/2协议之上，对TLS提供了很好的支持。当不需要证书认证时,可通过grpc.WithInsecure()选项跳过了对服务器证书的验证，没有启用证书的gRPC服务和客户端进行的是明文通信，
信息面临被任何第三方监听的风险。为了保证gRPC通信不被第三方监听、篡改或伪造，可以对服务器启动TLS加密特性

go 从 1.15版本开始推荐使用 SAN证书，如果使用不通过SAN生成的证书，会报以下错误信息。

```shell
rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0"
```

## 什么是SAN？

SAN(Subject Alternative Name) 是 SSL 标准 x509 中定义的一个扩展。使用了 SAN 字段的 SSL 证书，可以扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析。

## 证书生成流程

### 生成CA证书

```shell
# 生成.key  私钥文件
openssl genrsa -out ca.key 2048

# 生成.csr 证书签名请求文件
openssl req -new -key ca.key -out ca.csr  -subj "/C=GB/L=Shenzhen/O=github/CN=www.mark8s.com"

# 自签名生成.crt 证书文件
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt  -subj "/C=GB/L=Shenzhen/O=github/CN=www.mark8s.com"
```

- C=GB: C代表的是国家名称代码。
- L=Shenzhen: 代表地方名称,例如城市。
- O=github: 代表组织单位名称。
- CN=:=www.mark8s.com 代表关联的域名。

### 生成服务端证书

和生成 CA证书类似，不过最后一步由 CA 证书进行签名，而不是自签名。

先准备 openssl.cnf 文件，复制到当前目录，不同操作系统它存在的地址会不一样。

- linux系统 : /etc/pki/tls/openssl.cnf
- Mac系统: /System/Library/OpenSSL/openssl.cnf
- Windows：安装目录下 openssl.cfg 比如 D:\Program Files\OpenSSL-Win64\bin\openssl.cfg

复制到当前目录
```shell
$ cp /etc/pki/tls/openssl.cnf .
$ ls
openssl.cnf
```

开始生成证书
```shell
openssl genrsa -out server.key 2048

openssl req -new -key server.key -out server.csr \
	-subj "/C=GB/L=Shenzhen/O=github/CN=www.mark8s.com" \
	-reqexts SAN \
	-config <(cat openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.mark8s.com"))

openssl x509 -req -days 3650 \
   -in server.csr -out server.crt \
   -CA ca.crt -CAkey ca.key -CAcreateserial \
   -extensions SAN \
   -extfile <(cat openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.mark8s.com"))
```

至此，我们看一下当前目录下有哪些文件：
```shell
$ ls
ca.crt  ca.csr  ca.key  ca.srl  openssl.cnf  server.crt  server.csr  server.key
```

会用到的有下面三个：

- ca.crt
- server.crt 
- server.key

## 使用证书

服务端：

```shell
	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	creds, err := credentials.NewServerTLSFromFile("../pkg/x509/server.crt", "../pkg/x509/server.key")
	if err != nil {
		log.Fatalf("Failed to generate TLS credentials %v", err)
	}
	// 新建gRPC服务器实例,并开启TLS认证
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	proto.RegisterSimpleServer(grpcServer, &SecurityServer{})
```

客户端：

```shell
	creds, err := credentials.NewClientTLSFromFile("../pkg/x509/ca.crt", "www.mark8s.com")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	// 连接服务器
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
```

以上只是弄了服务端的tls加密，双向tls认证可以查看 [gRPC(Go)教程(四)---通过SSL/TLS建立安全连接](https://www.lixueduan.com/posts/grpc/04-encryption-tls/)

## Token认证
客户端发请求时，添加Token到上下文context.Context中，服务器接收到请求，先从上下文中获取Token验证，验证通过才进行下一步处理。





## 参考

[RPC编程(六):gRPC中的TLS认证](http://liuqh.icu/2022/02/23/go/rpc/06-tls/)

[gRPC(Go)教程(四)---通过SSL/TLS建立安全连接](https://www.lixueduan.com/posts/grpc/04-encryption-tls/)



