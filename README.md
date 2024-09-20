# OneAuth

统一验证及应用管理服务

[Demo](https://oa.veypi.com)


## Auth

code用来兑换本身应用token和兑换其他应用付出权限token
oa code 可以用来生成其他code和刷新本身code


### 依赖库

```bash
docker run -dit --name=tsdb  -v /Users/veypi/test/vdb:/victoria-metrics-data -p 8428:8428 victoriametrics/victoria-metrics -search.latencyOffset=1s
nats-server -c ./script/nats.cfg
```


