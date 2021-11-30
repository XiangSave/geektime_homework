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
$ d build -t sample-http-server:1.0  .

$ d run -itd -p 8080:8080 -v /home/xxx/httpServer/configs:/root/configs -v /home/xxx/httpServer/logs:/root/logs sample-http-server:1.0
3358f6cf265af595ae5922bcf30f6cf93fc3927e7263f323444134376461a140

$ curl localhost:8080/healthz
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

## week eight
### question01
+ 编写 Kubernetes 部署脚本将 httpserver 部署到 kubernetes 集群，需支持：
  + 优雅启动
  + 优雅终止
  + 资源需求和QoS保证
  + 探活
  + 日常运维需求、日志等级
  + 配置和代码分离

### question02
+ 除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。
+ 来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧
+ 在第一部分的基础上提供更加完备的部署 spec，包括（不限于）
  + Service
  + Ingress
+ 可以考虑的细节
  + 如何确保整个应用的高可用
  + 如何通过证书保证 httpServer 的通讯安全

### answer
+ kubernetes yaml 
  + deployment: k8smanifests/nginx-ingress-deployment.yaml
  + conf：k8smanifests/sample-http-server-conf-comfigMap.yaml
  + service: k8smanifests/sample-http-server-svc.yaml
  + tls secret:k8smanifests/test-sample-http-server-tls-secret.yaml
  + ingress:k8smanifests/http-server-ingress.yaml

### note
+ 创建 configMap 存储配置文件

```bash
$ k create configmap sample-http-server-conf --from-file=configs/
```
+ 创建 tls 证书

```bash
$ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout test-sample-http-server.key -out test-sample-http-server.crt -subj "/CN=*/O=xiang"
$ cat test-sample-http-server.key |base64 -w 0
$ cat test-sample-http-server.crt |base64 -w 0
```
+ 安装 ingress-nginx(k8smanifests/ingress-nginx-deployment.yaml)
  + 下载官网 deployment
```bash
$ wget https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.0/deploy/static/provider/cloud/deploy.yaml
```
  + 更改 deployment image
  + 更改 ingress-nginx Service spec.type 为 NodePort
  + 配置 CA

```bash
  $ CA=$(kubectl -n ingress-nginx get secret ingress-nginx-admission -ojsonpath='{.data.ca}')
  $ kubectl patch validatingwebhookconfigurations ingress-nginx-admission --type='json' -p='[{"op": "add", "path": "/webhooks/0/clientConfig/caBundle", "value":"'$CA'"}]'
```
