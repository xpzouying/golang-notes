# README

Go 代码直接打包前端代码，使用 embed 特性。（需要 Go 1.16 以上能力支持）


## 1. 运行

### 1.1 前端代码

代码统一在 [./dist](./dist/) 目录。

### 1.2 后端代码

直接运行，会打包前端源码。

```bash
go run .
```

### 1.3 测试

浏览器打开地址： `http://dev.zy.local:9090/`，更换成自己的地址即可。



## 2. Refs:

https://v0x.nl/articles/portable-apps-go-nextjs
