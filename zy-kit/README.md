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

    - `endpoint`支持：需要定一个`endpoint`（端点）。它的作用是：接收一个请求，解析请求参数，返回响应。
    所以，会将`service`所有的业务公共的方法都封装，以`endpoint`的形式暴露，按照指定的`decode`对请求进行解析请求，
    并且按照指定的`encode`进行响应的编码。
    例如，能够将HTTP的请求，转换成json形式，转换成`service`能够处理的请求，并且以json响应返回回去。

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


## 3-middleware-2

对于框架来说，Middleware的支持是必不可少的。那么，需要考虑Middleware需要封装在哪一层（如：`transport`、`endpoint`）。

由于，Middleware是不受协议约束，无论是`http`还是`grpc`，都需要对服务的调用进行Middleware的封装，所以应该是在`transport`到`endpoint`之间，并且经过Middleware封装后，仍然保持原有的endpoint能力。

所以，大致会得到`Middleware`的定义：

```go
type Middleware func(Endpoint) Endpoint
```

以调用接口的耗时统计功能为例，作一个TimeUsed Middlware。

TimeUsedMiddleware的实现如下：

```go
func WithTimeUsed(next Endpoint) Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		defer func(begin time.Time) {
			log.Printf("time_used: %s", time.Since(begin))
		}(time.Now())

		return next(ctx, request)
	}
}
```

但是，上面有个问题，就是比如我们依赖其他的组件，比如我们依赖一个自定义的`log.Logger`组件，那么这种方式就不是太合适。


## 3-metrics

这一示例中，我们增加`[prometheus](https://github.com/prometheus/client_golang/)`组件，进行metrics统计。

首先，按照[官方示例simple/main.go](https://github.com/prometheus/client_golang/blob/master/examples/simple/main.go)
将`prometheus`默认的功能加入进来。

```go
func NewPromHTTPHandler() http.Handler {
	return promhttp.Handler()
}
```

在`main.go`中增加`/metrics`接口，

```go
http.Handle("/metrics", zykit.NewPromHTTPHandler())
```

运行服务，

```bash
cd ./3-metrics
go run .
```

发起请求，

```bash
curl -XPOST -d '{"s": "ZouYing"}'  http://localhost:8080/upper
# "ZOUYING"

curl -XPOST -d '{"s": "ZouYing"}'  http://localhost:8080/count
# 7
```

打开浏览器：`localhost:8080/metrics`，可以看到默认的统计数据。

```
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 7
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.17.4"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 2.444808e+06
...
...
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 2
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```

上面日志中，

- `# HELP` 注释该指标的用途
- `# TYPE` 说明该指标的类型


接下来，我们对统计的数据进行丰富。增加`Prometheus`其他的信息，[示例](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#hdr-A_Basic_Example)。

**增加接口调用次数**

Prometheus中，提供了2个`Counter`：

- [NewCounter](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#NewCounter) - 不带label。

- [NewCounterVec](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#NewCounterVec) - 带有label。


直接引入`prometheus.Counter`，增加简单接口计数。

```go
type CounterOpts = prometheus.CounterOpts

type Counter struct {
	core prometheus.Counter
}

func (c *Counter) Inc() {
	c.core.Inc()
}

func NewCounter(opts CounterOpts) *Counter {

	counter := prometheus.NewCounter(opts)

	prometheus.MustRegister(counter)
	return &Counter{counter}
}
```

我们定义自己的`Counter`，只需要`Inc()`方法即可，里面就是调用prometheus自带的Counter的Inc方法。

```go
type instrumentMiddleware struct {
	requestCounter zykit.Counter
	svc StringService
}

func (mw *instrumentMiddleware) Upper(s string) string {
	mw.requestCounter.Inc()
	return mw.svc.Upper(s)
}

func (mw *instrumentMiddleware) Count(s string) int {
	mw.requestCounter.Inc()
	return mw.svc.Count(s)
}
```

创造一个middleware - `instrumentMiddleware`，里面实现了Service interface的接口要求。


**运行**

运行服务端：

```bash
cd 3-metrics

go run .
# 2021/12/08 13:44:07 listen on :8080
```

请求：

```bash
curl -XPOST -d '{"s": "ZouYing"}'  http://localhost:8080/count
# 7

curl -XPOST -d '{"s": "ZouYing"}'  http://localhost:8080/upper
# "ZOUYING"
```

打开浏览器，[http://localhost:8080/metrics](http://localhost:8080/metrics)，
查看自定义的统计：`request_count`。

```
# ...
# ...
# HELP request_count Counter of request
# TYPE request_count counter
request_count 2
```


## 4-stringsvc-multiple-transport

同时支持grpc和http协议。

安装`proto`解析：

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

安装`protobuf`，Macos系统：

```bash
brew install protobuf
```

> 如果仍然找不到，可能是命令`protoc-gen-go`找不到。
> 那么：`export PATH=$PATH:$HOME/go/bin`

生成`pb.go`文件：

```bash
protoc --go_out=. ./svc.proto
```
