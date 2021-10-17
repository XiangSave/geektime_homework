# sample http server

## week two

+ 编写一个 HTTP 服务器
  + 接收客户端 request，并将 request 中带的 header 写入 response header
  + 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
  + Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
  + 当访问 localhost/healthz 时，应返回200

+ main 为 cmd/httpServer/main.go

## week three

+ 构建本地镜像
+ 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
  + 位于 Dockerfile 中，使用多阶段构建
+ 将镜像推送至 Docker 官方镜像仓库。
+ 通过 Docker 命令本地启动 httpserver。

```bash
$ d build -t http_server:1.0  .

$ d run -itd --rm -p 80:80 http_server:1.0
3358f6cf265af595ae5922bcf30f6cf93fc3927e7263f323444134376461a140
$ curl localhost:80/healthz
health
```

+ 通过 nsenter 进入容器查看 IP 配置。

```bash
$ d inspect e6071393a5e9 |grep pid -i
            "Pid": 14748,
            "PidMode": "",
            "PidsLimit": null,
$ s nsenter  -t 14748 -n ip a
[sudo] password for xxx:
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: sit0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/sit 0.0.0.0 brd 0.0.0.0
94: eth0@if95: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```
