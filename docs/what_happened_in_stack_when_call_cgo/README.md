# cgo调用c函数时，调用栈发生了什么？

## 前言

分别有2个程序，完成相似的功能，其中一个是纯Go函数，另外一个是Go调用C函数，对于函数调用栈有什么区别？



**1、cgo程序**

```go
package main

/*
int sum(int a, int b) { return a+b; }
*/
import "C"

func main() {
	a, b := C.int(1), C.int(2)
	C.sum(a, b)
}
```



**2、Go程序**

```go
package main

func sum(a, b int) int {
	return a + b
}

func main() {
	a, b := 1, 2
	sum(a, b)
}
```
