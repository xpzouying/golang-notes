# channel的底层实现

## 基本概念

**有2种类型的channel**

1. 有缓存的channel，其定义为：
   - `make(chan int, 100)`
   - `make(chan struct{}, 10)`
2. 无缓存的channel，其定义为：
   - `make(chan int)`
   - `make(chan int, 0)`

**对channel的操作**

1. 向channel发送数据
   - `ch1 <- 1`
   - `ch2 <- struct{}{}`
2. 从channel获取数据
   - `v, ok := <-ch1`

     2个接收者。
     第一个获得接收的数据，第二个表示当前数值是否合法。
     若channel已被关闭，则第一个获取的值永远是0、false、nil等值，同时第二个值为false。

   - `v := <-ch2`
     
     1个接收者。
     获得的数据，如果channel已经关闭，会一直从channel中获取到0值（或false、nil等）

3. 关闭channel
   - `close ch1`

     关闭channel后，若有接收者，则一直会获取0值。

## 分析大纲

源码版本：

```bash
root@zy-dev:/usr/local/go/src/runtime# go version
go version go1.5.4 linux/amd64
```

分析channel底层实现时，也按照上层的操作逐一拆解，划分为：

1. 有缓存的channel
2. 无缓存的channel

分别分析对应的操作：

1. 创建
2. 发送
3. 接收
4. 关闭

## channel结构体

channel的主要思路是：

1. 使用一个buf数组保存channel中的数据；
   1. buf的大小为make时指定的大小，为dataqsize个。
   2. 由于buf的元素类型都一致（为elemtype），也即元素的字节大小一致，每个元素为elemsize字节大小，所以buf的内存空间也确定，为elemsize*dataqsize。 
   3. 循环复用buf数组，达到一种循环队列的效果。为了达到这种效果，借助2个成员标示队列的当前的数据头和尾。sendx为发送的index，recvx为接收的index。
   4. 当前buf中已有元素个数为qcount。
2. recvq、sendq保存当前被阻塞的接收和发送的goroutine。我们对于channel的发送和接收操作都会产生阻塞，当阻塞发生时，被阻塞的goroutine会先被保存起来，等待后续被唤醒。
   1. 使用队列进行保存。队列结构体为`waitq`结构体。内部实现使用了无锁并发队列的方式。


```go
type hchan struct {
        qcount   uint           // total data in the queue
        dataqsiz uint           // size of the circular queue
        buf      unsafe.Pointer // points to an array of dataqsiz elements
        elemsize uint16
        closed   uint32
        elemtype *_type // element type
        sendx    uint   // send index
        recvx    uint   // receive index
        recvq    waitq  // list of recv waiters
        sendq    waitq  // list of send waiters
        lock     mutex
}

type waitq struct {
        first *sudog
        last  *sudog
}
```


## 创建channel

根据前面提到的channel的创建形式主要分为有缓存的和无缓存的channel，对应的创建形式为，


```go
// filename: make_chan.go
package main

func main() {
	c1 := make(chan int, 10)
	c2 := make(chan int)
	_ = c1
	_ = c2
}
```

通过命令生成Golang汇编，

```bash
go build -gcflags "-N -l -S" make_chan.go &> make_chan.s
```

删除无关代码，得到对应的汇编代码为：

有缓存的buffer创建汇编为，

```
# make(chan int, 10)
# make的第一个参数为chan int
0x001d 00029 (.../make_chan.go:4)	LEAQ	type.chan int(SB), AX
0x0024 00036 (.../make_chan.go:4)	MOVQ	AX, (SP)
# make的第二个参数为10
0x0028 00040 (.../make_chan.go:4)	MOVQ	$10, 8(SP)
# 调用runtime中的makechan函数
0x0031 00049 (.../make_chan.go:4)	CALL	runtime.makechan(SB)
```

无缓存的buffer创建汇编为，

```
# make(chan int)
# make的第一个参数为chan int
0x0040 00064 (.../make_chan.go:5)	LEAQ	type.chan int(SB), AX
0x0047 00071 (.../make_chan.go:5)	MOVQ	AX, (SP)
# make的第二个参数为0
0x004b 00075 (.../make_chan.go:5)	MOVQ	$0, 8(SP)
# 调用runtime中的makechan函数
0x0054 00084 (.../make_chan.go:5)	CALL	runtime.makechan(SB)
```

由此可见当我们在make channel时调用的是同一个函数，只是第二个参数不同而已，unbuffer channel其实就是`make(chan int, 0)`。


**runtime.makechan函数分析**

源码见：`runtime/chan.go`，点击[github.com/go1.5](https://github.com/golang/go/blob/release-branch.go1.5/src/runtime/chan.go)。

代码删掉不必要的代码段。

```go
func makechan(t *chantype, size int64) *hchan {
	elem := t.elem

   // ...
   // 省略一些条件判断

	var c *hchan
	if elem.kind&kindNoPointers != 0 || size == 0 {
      // 如果是unbuffer channel，则分配hchan大小的内存空间。
		c = (*hchan)(mallocgc(hchanSize+uintptr(size)*uintptr(elem.size), nil, flagNoScan))
		if size > 0 && elem.size != 0 {
         c.buf = add(unsafe.Pointer(c), hchanSize)
		} else {
         // 并让hchan中的成员buf指向内存地址自身。（该地址提供给同步操作）
			c.buf = unsafe.Pointer(c)
		}
	} else {
      c = new(hchan)
      // buf指向环形buffer
		c.buf = newarray(elem, uintptr(size))
   }
   // 元素的个数
   c.elemsize = uint16(elem.size)
   // 元素的类型
   c.elemtype = elem
   // 环形buffer的大小
	c.dataqsiz = uint(size)

	return c
}
```
