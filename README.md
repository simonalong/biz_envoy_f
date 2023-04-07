### biz-f-service


记得将Dockerfile文件内容修改为对应的 Dockerfile-none-envoy 和 Dockerfile-have-envoy
```shell
# 有envoy
docker build -t biz-f-service:1.0.0-have-envoy .
# 有envoy【本地调试版本】
docker build -t biz-f-service:1.0.0-have-envoy-debug .
# 有envoy【链路数据】
docker build -t biz-f-service:1.0.0-have-envoy-trace .
# 没有envoy
docker build -t biz-f-service:1.0.0-none-envoy .
```

```shell
# 有envoy
docker run -d --name biz-f-service -p 18005:18005 biz-f-service:1.0.0-have-envoy
# 没有envoy
docker run -d --name biz-f-service -p 18005:18005 biz-f-service:1.0.0-none-envoy
```


```shell
# 测试
curl http://localhost:18015/api/f/cf/true
```

```shell
# 打镜像
docker build -t biz-f-service:1.0.0-none-envoy .
# 镜像保存
docker save biz-f-service:1.0.0-none-envoy -o biz-f-service.tar
# 镜像文件上传
scp -v biz-f-service.tar root@10.30.30.78:/root/zhouzy/biz-f-service.tar
# 开发环境中载入
docker load -i /root/zhouzy/biz-f-service.tar
```
