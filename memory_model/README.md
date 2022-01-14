# The Go Memory Model

## 指令重排

下列示例中，有可能会出现意想不到的情况。

```go
var a, b int

// goroutine A
go func() {
    a = 5
    b = 1
}()

// goroutine B
go func() {
    for b == 1 {}
    fmt.Println(a)
}()
```

对于goroutine B的输出，有可能会出现输出为：0。

**原因**

编译器和处理器在不改变语言规范所定义的行为前提下，有可能会进行：`指令重排`的操作。

以`goroutine A`为例，

```go
a = 5
b = 1
```

在同一个goroutine中，先对a赋值、还是对于b赋值，实际上是一样的，因为不会影响Go的任何语言规范。

所以，有可能会出现一种奇怪的现象，就是：经过指令重排后，b=1被放在前面，从而会导致a=5这个操作没有做的时候，goroutine B又开始运行，这样会输出0。

如果要约定[Go内存模型](https://go.dev/ref/mem)中提到的各种同步机制。


**验证**

```go
package main

import "fmt"

func do() {

	var a, b int

	go func() {
		a = 5
		b = 1
	}()

	go func() {
		for b == 1 {
		}

		if a == 0 {
			fmt.Println(a)
		}
	}()
}

func main() {

	for {
		do()
	}
}
```


