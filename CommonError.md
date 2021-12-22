# 常见错误


## 以协程 goroutines 运行的闭包会发生什么

见下列示例

```go
package main

import (
        "fmt"
        "time"
)

func main() {
        values := []int{1, 2, 3}

        for _, v := range values {
             go func() {
                     fmt.Println(v)
             }()
        }

        time.Sleep(1 * time.Second)
}
```

结果会输出：`3 3 3`。

为什么呢？

由于在Go中是使用`:=`可以定义一个变量，那么相当于我们在for循环里面一直在使用同一个变量：`v`，然后在goroutine中一直复用。

类似于下列代码，

```go
package main

import (
        "fmt"
        "time"
)

func main() {
        values := []int{1, 2, 3}

        var v int
        for _, v = range values {
                go func() {
                        fmt.Println(v)
                }()
        }
        time.Sleep(1 * time.Second)
}
```

如果我们需要修复，那么就需要让每个goroutine都有自己的`v`变量，可以下列两种方式：

1. 局部变量；

2. 函数参数传递；

```go
package main

import (
        "fmt"
        "time"
)

func main() {
        values := []int{1, 2, 3}

        // 方式1 - 让每个函数都有自己的局部变量
        for _, v := range values {
                // 重新定义局部变量
                v := v

                go func() {
                        fmt.Println(v)
                }()
        }

        // 方式2 - 传递函数参数
        for _, v := range values {
                // 重新定义局部变量

                go func(n int) {
                        fmt.Println(n)
                }(v)
        }

        time.Sleep(1 * time.Second)
}
```

