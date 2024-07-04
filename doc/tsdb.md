# TSBD

```bash

docker run -dit --name=tsdb  -v /Users/veypi/test/vdb:/victoria-metrics-data -p 8428:8428 victoriametrics/victoria-metrics -search.latencyOffset=1s -retentionPeriod=10y -search.cacheTimestampOffset=1h


# 查询所有数据
curl http://localhost:8428/prometheus/api/v1/series -d 'match[]={__name__=~".*"}'

# 删除所有数据
curl http://localhost:8428/prometheus/api/v1/series\?start=1593100800\&end=1719308095\&step=2d22h40m50s -d 'match[]={__name__=~".*"}'
curl http://localhost:8428/api/v1/admin/tsdb/delete_series -d 'match[]={__name__=~".*"}'

# 查询所有标签
curl http://localhost:8428/prometheus/api/v1/labels

#  插入数据
curl -X POST http://localhost:8428/api/v1/import/prometheus -d 'abd{foo="bar",a="15"} 25' 
curl -X POST http://localhost:8428/api/v1/import -d '{"metric":{"__name__":"foo","job":"node_exporter"},"values":[4],"timestamps":[1719819673000]}'

curl -G 'http://localhost:8428/api/v1/export' -d 'match={__name__="abd"}'
curl -G 'http://localhost:8428/api/v1/export' -d 'match[]={__name__="abd"}'


```
