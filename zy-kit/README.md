# zy-kit

Yet another [go-kit](https://github.com/go-kit/kit).

`go-kit`是一套微服务框架。通过这套框架我们能快速在我们的服务内部集成各种微服务的组件，例如：限流、Logging+Metrics+Trace等等。

该系列分为几步实现一套微服务的框架，

1. 先构建好一个web服务框架。能把服务先跑起来。

2. 集成通用的日志中间件（Middleware）。

3. 加入Prometheus，进行metrics统计。

4. 加入除了`HTTP`协议以外的其他协议支持。


## 1-kit-base

**参考资料：**

- [示例1 - stringsvc](https://gokit.io/examples/stringsvc.html)


框架的能力和用户业务的能力，拆分如下：

1. 框架能力：

    - `endpoint`支持：需要定一个`endpoint`（端点）。它的作用是，将`service`所有的API都封装，以`endpoint`的形式暴露。
能够将HTTP的请求，转换成`service`能够处理的请求。

    - `transport`支持：有了上面的`endpoint`，就需要定义`transport`，能够支撑我们对应的协议，比如：`http`支持、或者`grpc`等协议的支持。
这里的主要作用是：（以http为例），能够提供`http`对应的服务、能够处理http请求、支持http请求到`endpoint`处理。

1. 用户业务能力：

    - `interface Service` - 定义服务的行为。

    - `struct service` - 服务interface的实现


**运行**

运行`1-kit-base`中的代码，

```bash
cd zy-kit/1-kit-base/

go run .
# 2021/12/06 13:10:59 listen on :8080
```

请求服务：

```bash
curl -XPOST -d '{"s": "ZouYing"}'  http://localhost:8080/upper
# "ZOUYING"
```


## 2-middleware

非框架实现的方式。

我们创建一个`struct`，里面实现`Service interface`的方法。

```go
type TimeUsedMiddleware struct {
	next StringService
}

func (mw *TimeUsedMiddleware) Upper(s string) (res string) {
	defer func(begin time.Time) {
		log.Printf("time_used: %s", time.Since(begin))

	}(time.Now())

	return mw.next.Upper(s)
}
```

调用时，`svr`变量一直时`Service interface`的类型，所以我们在封装时，只需要实现接口对应的方法即可。


```go
var svr StringService
{
    // 核心的service
    svr = &simpleStringServer{}

    // 封装中间件
    svr = WithTimeUsedMiddleware(svr)
}
```

详细的示例代码在：`2-middleware`。

**运行**

```bash
cd 2-middleware

go run .
# 2021/12/06 17:16:13 listen on :8080
```

**请求**

```bash
curl -XPOST -d '{"s": "ZouYing"}'  http://localhost:8080/upper
# "ZOUYING"
```

同时，服务端会输出请求的耗时：

```bash
2021/12/06 17:16:35 simple_string_server: got=ZouYing result=ZOUYING
2021/12/06 17:16:35 time_used: 48.25µs
```
