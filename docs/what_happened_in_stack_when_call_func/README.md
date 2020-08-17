# 调用Go函数时，内存栈发生了什么？

在Go中，函数调用发生时，内存调用栈都发生了什么？

## 示例

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

**通过命令生成汇编代码：**

```bash
# 生成 main.o
go tool compile -N -l main.go

go tool objdump main.o
```

**输出汇编内容：**

```
TEXT %22%22.sum(SB) gofile../Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go
  main.go:3             0x3d9                   48c744241800000000      MOVQ $0x0, 0x18(SP)
  main.go:4             0x3e2                   488b442408              MOVQ 0x8(SP), AX
  main.go:4             0x3e7                   4803442410              ADDQ 0x10(SP), AX
  main.go:4             0x3ec                   4889442418              MOVQ AX, 0x18(SP)
  main.go:4             0x3f1                   c3                      RET

TEXT %22%22.main(SB) gofile../Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go
  main.go:7             0x406                   65488b0c2500000000      MOVQ GS:0, CX           [5:9]R_TLS_LE
  main.go:7             0x40f                   483b6110                CMPQ 0x10(CX), SP
  main.go:7             0x413                   7641                    JBE 0x456
  main.go:7             0x415                   4883ec30                SUBQ $0x30, SP
  main.go:7             0x419                   48896c2428              MOVQ BP, 0x28(SP)
  main.go:7             0x41e                   488d6c2428              LEAQ 0x28(SP), BP
  main.go:8             0x423                   48c744242001000000      MOVQ $0x1, 0x20(SP)
  main.go:8             0x42c                   48c744241802000000      MOVQ $0x2, 0x18(SP)
  main.go:9             0x435                   488b442420              MOVQ 0x20(SP), AX
  main.go:9             0x43a                   48890424                MOVQ AX, 0(SP)
  main.go:9             0x43e                   48c744240802000000      MOVQ $0x2, 0x8(SP)
  main.go:9             0x447                   e800000000              CALL 0x44c              [1:5]R_CALL:%22%22.sum
  main.go:10            0x44c                   488b6c2428              MOVQ 0x28(SP), BP
  main.go:10            0x451                   4883c430                ADDQ $0x30, SP
  main.go:10            0x455                   c3                      RET
  main.go:7             0x456                   e800000000              CALL 0x45b              [1:5]R_CALL:runtime.morestack_noctxt
  main.go:7             0x45b                   eba9                    JMP %22%22.main(SB)
```

