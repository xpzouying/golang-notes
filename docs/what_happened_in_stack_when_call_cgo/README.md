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



## 分析

### 生成汇编

使用命令生成汇编：

```bash
# 方法1
go tool compile -S main_go.go

# 方法2
go build -gcflags -S main_cgo.go                                             
```



**1、生成Go代码：`go build -gcflags -S main_go.go`**

```
"".sum STEXT nosplit size=19 args=0x18 locals=0x0
    main_go.go:3)   TEXT    "".sum(SB), NOSPLIT|ABIInternal, $0-24
    main_go.go:3)   PCDATA  $0, $-2
    main_go.go:3)   PCDATA  $1, $-2
    main_go.go:3)   FUNCDATA    $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    main_go.go:3)   FUNCDATA    $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    main_go.go:3)   FUNCDATA    $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    main_go.go:4)   PCDATA  $0, $0
    main_go.go:4)   PCDATA  $1, $0
    main_go.go:4)   MOVQ    "".b+16(SP), AX
    main_go.go:4)   MOVQ    "".a+8(SP), CX
    main_go.go:4)   ADDQ    CX, AX
    main_go.go:4)   MOVQ    AX, "".~r2+24(SP)    
    main_go.go:4)   RET    

"".main STEXT nosplit size=1 args=0x0 locals=0x0
    main_go.go:7)   TEXT    "".main(SB), NOSPLIT|ABIInternal, $0-0    
    main_go.go:7)   PCDATA  $0, $-2    
    main_go.go:7)   PCDATA  $1, $-2    
    main_go.go:7)   FUNCDATA    $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)    
    main_go.go:7)   FUNCDATA    $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)    
    main_go.go:7)   FUNCDATA    $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)    
    main_go.go:9)   PCDATA  $0, $-1    
    main_go.go:9)   PCDATA  $1, $-1    
    main_go.go:9)   RET    
```



**2、生成CGO代码：``go build -gcflags -S main_cgo.go`**

```

```

