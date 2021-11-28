# HTTP服务如何优雅退出和重启

后端服务常常有优雅退出和重启的需求。

分2部分实现：

1. 实现优雅退出。

2. 结合现实的场景，进行优雅重启。

## 优雅退出

### 正常的HTTP Server

为了手动模拟现实情况，在HTTP处理函数中插入了3秒接口处理耗时。

代码在：`graceful_shutdown/normal_server/`目录下。

```go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)

		fmt.Fprintf(w, "hello world")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

运行程序：

```bash
# run server
go run .
```

请求服务：

```bash
curl http://localhost:8080
# hello world
```

### 问题

当服务正在处理请求时，当前请求没有结束，直接退出服务会导致客户端接收到下列问题。

```bash
curl http://localhost:8080
# curl: (52) Empty reply from server
```

在发出请求后，立刻在启动服务的命令中按下`ctrl+c`退出。

启动服务的命令会接收到`interrupt`系统信号通知，然后退出。

```bash
go run .
# signal: interrupt
```

### 思路

我们希望把正在进行中的请求处理完后，再退出。那么，

1. 需要能够捕获系统信号通知。

2. 让HTTP Server能够优雅退出。借助[Server.Shutdown](https://pkg.go.dev/net/http#Server.Shutdown)。


### 改造

借助上面的思路，需要进行下列改造。

1. 捕获信号：使用[os.Signal](https://pkg.go.dev/os/signal#Notify)监听interrupt信号（`os.Interrupt`）。

2. 在原有的代码中`http.ListenAndServe`中，会创建一个私有的server，我们无法直接调用对应的`Shutdown`进行退出。
所以需要创建一个新的、显示的Server。

```go
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

改动后的代码在`./graceful_shutdown/graceful_shutdown/main.go`。

代码解读：

```go
func main() {
	// 定义需要捕获的信号，并且通过channel进行传递。
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // os.Interrupt = syscall.SIGINT

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Fprintf(w, "hello world")
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	go func() {
		log.Println(server.ListenAndServe())
	}()

	// 等待接收信号
	s := <-c
	log.Printf("receive signal: %v", s)

	log.Println("http server shutdown: ", server.Shutdown(context.Background()))
}
```

1. 监听`ctrl+c`的系统信号。`os.Interrupt`也就是`syscall.SIGINT`信号。
2. 显示的创建`http.Server`，当接收退出信号，调用`server.Shutdown`进行优雅退出。

[Shutdown函数](https://pkg.go.dev/net/http#Server.Shutdown)在Go1.8版本添加进来。
主要作用是不中断所有正在活动的连接。具体的过程是：

1. 关闭所有的监听（Listener）。

2. 关闭所有的空闲连接（idle connection）。

3. 等待所有的连接变为空闲（idle）状态。

4. 关闭服务。

点击展开[官方示例](https://pkg.go.dev/net/http#example-Server.Shutdown)。

> 注意⚠️
>
> `log.Println(server.ListenAndServe())`中日志输出函数不能换成：`log.Fatal`。
> 否则，会立刻退出。目前还不知道为什么会导致这个问题。看了Fatal函数的实现，只有一个差别，
> 多了个`os.Exit(1)`。
>
> 猜测：在`os.Exit(1)`中，可能触发了强制退出的某种行为。
