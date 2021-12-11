# sample http server

## week ten
### question
+ 为 HTTPServer 添加 0-2 秒的随机延时
+ 为 HTTPServer 项目添加延时 Metric
+ 将 HTTPServer 部署至测试集群，并完成 Prometheus 配置
+ 从 Promethus 界面中查询延时指标数据
+ 创建一个 Grafana Dashboard 展现延时分配情况

### answer
+ metrics:pkg/metrics


### note
+ 测试 metrics

```bash
# 发送请求
$ while True;do ;curl localhost:8080/aaa & ;done

# 查看结果
$ watch -n 1 "curl localhost:8080/metrics |grep httpserver_execution_used_second_bucket"
```