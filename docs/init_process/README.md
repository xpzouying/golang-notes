# Go程序的初始化工作


对于用户程序来说，main package中的main()是程序的入口函数。

```go
package main

func main() {
	println("hello world")
}
```

思考如下问题，

- main()是如何被加载启动？
- 在main()运行前，还做了哪一些初始化工作？