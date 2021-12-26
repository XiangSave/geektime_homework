# sample http server

## week twelve
### question
+ 把我们的 httpserver 服务以 istio ingress gateway 的形式发布出来。以下是你需要考虑的几点：
  + 如何实现安全保证
  + 七层路由规则
  + 考虑 open tracing 的接入

### answer
#### 使用 istio ingresss gateway 将 httpServer 发布出来，并配置证书，使用 https 访问
+ yaml：
  + k8smainifests/http-server-01-deployment.yaml
+ 配置
  + default namespace 注入 istio sidecar

  ```bash
  $ kubectl label ns default istio-injection=enabled
  namespace/default labeled
  ```
  + 申请、配置证书
  
  ```bash
  # 签发证书
  $ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout http-server.key -out http-server.crt -subj "/CN=*/O=xiang"
  # 将证书配置在 istio 上
  $ kubectl create -n istio-system secret tls http-server-credential --key=http-server.key --cert=http-server.crt
  ```