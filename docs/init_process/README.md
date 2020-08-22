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


-----

## 分析过程

### 本地环境

- 操作系统：Debian GNU/Linux 8 (jessie)
- Go版本：go version go1.5.4 linux/amd64



首先需要找到整个程序的[entry point](https://en.wikipedia.org/wiki/Entry_point)，然后再梳理对应的流程。



### entry point

查看程序的`entry point`有多个方法，

- 使用`readelf`查看
- 使用调试工具
  - [go-delve/delve](https://github.com/go-delve/delve)
  - gdb



**1、使用`readelf`读取程序的相关信息：**

a. 编译源码：`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags "all=-N -l" main.go`

b. 使用`readelf -h main`查看相关信息：

  ```bash
# readelf -h main
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              EXEC (Executable file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x44e3b0
  Start of program headers:          64 (bytes into file)
  Start of section headers:          456 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         7
  Size of section headers:           64 (bytes)
  Number of section headers:         21
  Section header string table index: 6
  ```

其中，`Entry point address`为真正的入口地址。



使用`objdump -S --start-address=0x44e3b0 | more`查看入口函数的指令。

```bash
./main:     file format elf64-x86-64


Disassembly of section .text:

000000000044e3b0 <_rt0_amd64_linux>:
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT _rt0_amd64_linux(SB),NOSPLIT,$-8
        LEAQ    8(SP), SI // argv
  44e3b0:       48 8d 74 24 08          lea    0x8(%rsp),%rsi
        MOVQ    0(SP), DI // argc
  44e3b5:       48 8b 3c 24             mov    (%rsp),%rdi
        MOVQ    $main(SB), AX
  44e3b9:       48 8d 05 10 00 00 00    lea    0x10(%rip),%rax        # 44e3d0 <main>
        JMP     AX
  44e3c0:       ff e0                   jmpq   *%rax
        ...
```



**2、使用`dlv`查看入口**

或者使用`dlv debug main.go`打开调试，

```bash
# dlv debug main.go

Type 'help' for list of commands.
(dlv) l
> _rt0_amd64_linux() /usr/local/go/src/runtime/rt0_linux_amd64.s:8 (PC: 0x465660)
Warning: debugging optimized function
     3: // license that can be found in the LICENSE file.
     4:
     5: #include "textflag.h"
     6:
     7: TEXT _rt0_amd64_linux(SB),NOSPLIT,$-8
=>   8:         JMP     _rt0_amd64(SB)
     9:
    10: TEXT _rt0_amd64_linux_lib(SB),NOSPLIT,$0
    11:         JMP     _rt0_amd64_lib(SB)
```

其中`rt0`是`runtime0`的缩写。

-----

### _rt0_amd64_linux入口

_rt0_amd64_linux位于`rt0_linux_amd64.s`文件中。

```
TEXT _rt0_amd64_linux(SB),NOSPLIT,$-8
	LEAQ	8(SP), SI // argv
	MOVQ	0(SP), DI // argc
	MOVQ	$main(SB), AX
	JMP	AX
	
TEXT main(SB),NOSPLIT,$-8
	MOVQ	$runtime·rt0_go(SB), AX
	JMP	AX
```

首先，把`argv`数组的地址放在SI寄存器，`argc`参数值放在DI寄存器中，调用`$main`方法。

`$main`方法又调用了`$runtime.rt0_go`方法，该方法内部实际做了初始化的工作。

![image-20200822230551117](image-20200822230551117.png)



### rt0_go初始化工作

```
TEXT runtime·rt0_go(SB),NOSPLIT,$0
	// copy arguments forward on an even stack
	MOVQ	DI, AX		// argc
	MOVQ	SI, BX		// argv
	SUBQ	$(4*8+7), SP		// 2args 2auto
	ANDQ	$~15, SP  // 栈顶寄存器SP 按照16字节对齐
	MOVQ	AX, 16(SP)  // 保存寄存器AX值：argc值
	MOVQ	BX, 24(SP)  // 保存寄存器BX值：argv地址
	
	// 初始化g0
	// create istack out of the given (operating system) stack.
	// _cgo_init may update stackguard.
	MOVQ	$runtime·g0(SB), DI  // g0的地址放入DI寄存器
	LEAQ	(-64*1024+104)(SP), BX  // BX寄存器的值：SP+(-64*1024+104)
	MOVQ	BX, g_stackguard0(DI)  // g0.stackguard0 = BX的值（即SP+(-64*1024+104)）
	MOVQ	BX, g_stackguard1(DI)  // g0.stackguard1 = BX的值（即SP+(-64*1024+104)）
	MOVQ	BX, (g_stack+stack_lo)(DI)  // g0.stack.lo = BX的值（即SP+(-64*1024+104)）
	MOVQ	SP, (g_stack+stack_hi)(DI)  // g0.stack.hi = SP
```



初始化完后，g0与内存栈的状态如下：

![image-20200822233429873](image-20200822233429873.png)

