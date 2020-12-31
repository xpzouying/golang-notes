# Prometheus + Grafana

## 启动Prometheus


参考资料：

- [Minimal Prometheus setup](https://dev.to/ablx/minimal-prometheus-setup-with-docker-compose-56mp)



```bash
docker-compose up -d
```

启动后，测试。

```bash
curl -X POST http://localhost:9000/-/reload
```


```bash
curl http://localhost:9090/api/v1/label/job/values
```