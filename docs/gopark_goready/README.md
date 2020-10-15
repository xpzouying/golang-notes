# GOPARK与GOREADY

## 1、前言

当我们使用Go编程时，程序并发运行时，对于操作系统仍是线程在运行程序指令，但是对于Go层面，加入了`协程（goroutine）`的机制。

所以线程运行指令就变成了协程运行，从而有了协程运行的基本流程是（main函数也对应着协程）：

1. 创建新的协程Goroutine
2. 将该协程放在等待队列中
3. 等待合适的调度时机，被调度运行

由于程序中，不止有少数的几个协程在运行，所以就会出现协程调度，其主要作用是：

1. 找到合适的goroutine，让其运行起来
2. 在合适的时间点，把正在运行的goroutine切换下来

有上面的点，引出了：

1. Go是如何将现在在运行的goroutine切换下来；
2. Go是如何将goroutine调度运行起来的；


## 2、停止goroutine

在分析[channel的底层实现](https://github.com/xpzouying/golang-notes/issues/17)时，可以看到当协程阻塞时会调用`gopark`。


**2.1、看源码前的思考**

先记录一些自己的猜测和想法。

在把正在运行的goroutine切换下来时，我们需要做下列一些工作：

1. 解绑g和m。g只有与m绑定时，才能运行。真正在执行代码指令的仍然是m，所以需要把当前运行的g和m进行解绑，好让下一个g与m进行绑定并运行。
2. 记录g当前的运行状态。当前g从队列中被重新调度起来后，能从之前暂时的指令继续恢复运行，所以需要记录几个状态：
   - 栈的信息
   - PC信息
3. 将当前g放置到对应的队列中，等待时机。



---

## 备注笔记

1. **问：** 协程调度器在调度（`schedule`）时，如何从一个Goroutine切换到另外一个Goroutine运行？

**答：** 

理论上来说，对于程序运行时来说，进程/线程/协程都是一段指令代码，程序运行的过程就是CPU顺序执行指令的过程。

CPU具体执行那一条指令，取决于`PC寄存器`，所以想要运行另一个Goroutine的话，只需要把`PC寄存器`设置为新的Goroutine的指令地址后即可。

那么，实际上Go是如何实现的呢？在[execute()](https://github.com/golang/go/blob/release-branch.go1.5/src/runtime/proc1.go#L1336)函数中，运行新的Goroutine时，会调用[gogo(&gp.sched)](https://github.com/golang/go/blob/release-branch.go1.5/src/runtime/proc1.go#L1380)函数做切换。

对于[gogo](https://github.com/golang/go/blob/release-branch.go1.5/src/runtime/asm_amd64.s#L158)函数是汇编实现。

```
TEXT runtime·gogo(SB), NOSPLIT, $0-8
   // 省略其他代码

   MOVQ	gobuf_pc(BX), BX  // 将上次保存的PC地址读取出来
   JMP	BX                // 跳转到上次保留的PC地址
```

通过更新`PC寄存器`，然后跳转到对应的地址，即可实现Goroutine的切换。