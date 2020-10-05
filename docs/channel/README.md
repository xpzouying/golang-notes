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


## 发送数据到channel

**先分析`unbuffered channel`**

```go
package main

func main() {
   c2 := make(chan int)
   c2 <- 1
}
```

导出汇编，获得`c2<-1`的对应的汇编代码为：

```bash
go build -gcflags "-N -l -S" make_chan.go &> make_chan.s
```

```
0x0036 00054 (channel/make_chan.go:4)	MOVQ	16(SP), AX
0x003b 00059 (channel/make_chan.go:4)	MOVQ	AX, "".c2+24(SP)
0x0040 00064 (channel/make_chan.go:5)	MOVQ	AX, (SP)
0x0044 00068 (channel/make_chan.go:5)	LEAQ	""..stmp_0(SB), AX
0x004b 00075 (channel/make_chan.go:5)	MOVQ	AX, 8(SP)

0x0050 00080 (channel/make_chan.go:5)	CALL	runtime.chansend1(SB)
```

可以看到调用了`runtime.chansend1`函数。

```go
func chansend1(t *chantype, c *hchan, elem unsafe.Pointer) {
   chansend(t, c, elem, true, getcallerpc(unsafe.Pointer(&t)))
}

func chansend(t *chantype, c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
   // ...

   // 如果往nil channel中发送，则抛出unreachable的异常
   if c == nil {
      if !block {
         return false
      }
      // code src: trace.go
      // traceEvGoStop = 16 // goroutine stops (like in select{}) [timestamp, stack]
      //
      // gopark函数会将goroutine变成waiting状态并调用unlockf。
      // 如果unlockf返回false，则goroutine会被唤醒。
      gopark(nil, nil, "chan send (nil chan)", traceEvGoStop, 2)
      throw("unreachable")
   }

   // ...

   lock(&c.lock)
   if c.closed != 0 {
      // 如果往已关闭的channel发送，则抛出异常：send on closed channel
      unlock(&c.lock)
      panic("send on closed channel")
   }

   if c.dataqsiz == 0 { // 同步channel
      // 从接收队列中寻找接收者
      sg := c.recvq.dequeue()
      if sg != nil {
         // 如果找到接收者
         unlock(&c.lock)

         recvg := sg.g
         if sg.elem != nil {
            // 如果接收者的data element（unsafe.Pointer类型）不为nil，
            // 则调用syncsend方法直接拷贝需要发送的元素到接收者中的sudo.elem成员。
            // 
            // TODO(zy): 在syncsend中，函数结束前会调用sg.elem = nil。
            // 也即在讲elem的数据拷贝到接收者的elem地址后，elem的地址会被置空，难道接收端有其他的地方记录该地址？
            syncsend(c, sg, ep)
         }
         // 协程对象g 被唤醒时的参数
         recvg.param = unsafe.Pointer(sg)

         // 使用goready唤醒接收者的协程
         goready(recvg, 3)
         return true
      }
      
      // ...

      // 如果没有接收者，则在当前channel阻塞
      gp := getg()
      
      // 生成sudog结构对象
      // 由于当前没有接收者，所以需要将当前的发送者给暂存起来，等待有接收者时被唤醒。
      // 当前阻塞的goroutine会被保存在sudog的结构体对象中
      mysg := acquireSudog()
      mysg.releasetime = 0
      
      // 保存发送的数据
      mysg.elem = ep
      // sudog的等待队列（waiting list）
      mysg.waitlink = nil
      // 发送者协程的等待队列（waiting list）
      // TODO(zy):该waiting和上面的waitlink的工作状态细节是什么？
      gp.waiting = mysg
      // 发送者goroutine作为sudog的成员
      mysg.g = gp
      mysg.selectdone = nil
      // 没有唤醒时的参数
      gp.param = nil

      // 入队到channel的发送者队列中
      c.sendq.enqueue(mysg)
      // 阻塞当前发送者goroutine。将当前的goroutine切换到waiting状态，并且释放c.lock锁。
      // traceEvGoBlockSend = 22 // goroutine blocks on chan send [timestamp, stack]
      goparkunlock(&c.lock, "chan send", traceEvGoBlockSend, 3)

      // zy: 上一行的goparkunlock会被阻塞住，除非被对应的goready唤醒。
      // 所以，当代码运行到这里时表示阻塞的发送者已经被唤醒，从而对应的资源需要被清理掉
      gp.waiting = nil
      if gp.param == nil {
         if c.closed == 0 {
            throw("chansend: spurious wakeup")
         }
         panic("send on closed channel")
      }
      gp.param = nil

      releaseSudog(mysg)
      return true
   }

   // 处理buffer channel
   // ...
}
```

总结unbuffered channel发送流程，对于下列语句：

```go
c <- 1
```

1. c为nil时，发送者的goroutine会被设置为waiting状态。
2. c为closed时，会panic错误：send on closed channel。
3. 尝试从c的recvq中获取一个接收者。若成功获取的话，则：
   1. 使用syncsend将数据拷贝给接收者，并
   2. 使用goready唤醒接收者
4. 若没有接收者，则需要将当前发送者的goroutine阻塞。具体流程：
   1. 将goroutine和数据打包成sudog结构体，放入sendq的队列中。
   2. 调用goparkunlock将当前的goroutine切换到waiting状态。


**分析`buffered channel`**

直接分析`chansend`函数，删除无关代码。

```go
func chansend(t *chantype, c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {

   // ...
   lock(&c.lock)

   // ... 省略处理unbuffered channel

   // 处理buffered channel，也称为异步channel

   // 如果channel的qcount>=dataqsize，即当前channel已有的数据>=长度，则一直循环判断，直到有空间可以存放。
   for futile := byte(0); c.qcount >= c.dataqsiz; futile = traceFutileWakeup {
      // 如果没有空间，则将发送者goroutine和发送元素封装成sudog结构体
      gp := getg()
      mysg := acquireSudog()
      mysg.releasetime = 0
      mysg.g = gp
      mysg.elem = nil
      mysg.selectdone = nil
      // 将发送者sudog放入发送者队列中
      c.sendq.enqueue(mysg)
      // 调用goparkunlock设置goroutine为waiting状态，等待后续被唤醒
      goparkunlock(&c.lock, "chan send", traceEvGoBlockSend|futile, 3)

      // 如果阻塞的goroutine被唤醒，则继续尝试for循环
      releaseSudog(mysg)
      lock(&c.lock)
      if c.closed != 0 {
         unlock(&c.lock)
         panic("send on closed channel")
      }
   }

   // 此时，可以进行发送。将发送数据写入到channel的buffer中
   typedmemmove(c.elemtype, chanbuf(c, c.sendx), ep)
   // 发送的index递增1
   c.sendx++
   // 如果超出buffer长度，则置0。
   if c.sendx == c.dataqsiz {
      c.sendx = 0
   }
   // 现存的元素个数也递增1
   c.qcount++

   // 尝试唤醒一个等待的接收者
   // 从channel的recvq获取一个接收者sg
   sg := c.recvq.dequeue()

   if sg != nil {
      // 如果成功获取到接收者，则使用goready唤醒
      recvg := sg.g
      unlock(&c.lock)
      // goready是将g的状态从waiting修改为runnable
      goready(recvg, 3)
   } else {
      // 如果没有获取到接收者，则直接结束
      unlock(&c.lock)
   }

   return true
}
```

总结：

1. 判断channel buffer是否有空间存放发送者的数据。
2. 如果没有空间，则：
   1. 封装发送者goroutine和数据为sudog结构体，并放入`sendq`发送者等待队列中。
   2. 调用`goparkunlock`，将发送者协程转换为`waiting`状态，等待唤醒尝试。
   3. 若当前阻塞的协程被其他协程唤醒，则再次判断是否有空间存放元素。若依旧没有空间，则重复2.1。若有空间，则继续步骤3。
3. 将发送的数据写入buffer，并：
   1. sendx累加1，表示发送元素的index后移。
   2. qcount累加1，表示buffer中的元素总数增加。
4. 尝试从`recvq`中唤醒一个接收者。
   1. 如果有接收者，则使用`goready`唤醒。结束。
   2. 如果没有接收者，直接结束。



## 接收channel数据

生成Go汇编，（在此不重复），`<-ch`语句会被转换为：`runtime/chan.go`中的函数：`func chanrecv(t *chantype, c *hchan, ep unsafe.Pointer, block bool) (selected, received bool)`，包含两个返回值：

1. 返回数据，true
2. 如果是被关闭的channel，则返回：0值（nil，0，false等等），false。

分析流程还是和接收数据的方式一致，分别按照unbuffered和buffered。

**1、从unbuffered channel获取数据**

```go
// block总是为true
func chanrecv(t *chantype, c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) {
   if c == nil {
      // 如果channel为nil，则goroutine会被设置为waiting状态，放入等待队列中。
      // 会被系统检测协程检测到永远不会被唤醒，从而导致deadlock的错误。
      gopark(nil, nil, "chan receive (nil chan)", traceEvGoStop, 2)
      throw("unreachable")
   }

   if !block && (c.dataqsiz == 0 && c.sendq.first == nil ||
      c.dataqsiz > 0 && atomicloaduint(&c.qcount) == 0) &&
      atomicload(&c.closed) == 0 {
      return
   }

   lock(&c.lock)  // 保护channel的状态
   if c.dataqsiz == 0 {  // unbuffered channel，即同步channel
      if c.closed != 0 {
         // 如果是从已经关闭（closed）的channel中获取数据，返回。第2个值总是0值。
         return recvclosed(c, ep)
      }

      // 从sendq（发送者阻塞队列）中获取一个发送者
      // 由于是unbuffered channel，数据不会保存在channel buffer中，
      // 所以如果需要获取数据的话，必然是从发送者的sudog中获取到。
      sg := c.sendq.dequeue()
      if sg != nil { // 1、如果获取的一个发送者
         unlock(&c.lock)

         if ep != nil {
            // 将发送者的数据拷贝出来
            typedmemmove(c.elemtype, ep, sg.elem)
         }
         
         // 置空发送者的elem。
         // （zy：为了发送者更好的GC，还是为了发送者做进一步的判断？）
         sg.elem = nil
         gp := sg.g
         gp.param = unsafe.Pointer(sg)  // 切换goroutine时的参数
         goready(gp, 3)  // 使用goready将发送者阻塞的协程置成runnable状态，等待重新被调度。
         selected = true
         received = true
         return
      }
      
      // 2、如果没有发送者。我们需要阻塞当前接收者协程。
      // zy：与前面发送者类似，会封装成sudog结构体，放入接收者阻塞队列中，即recvq中。
      gp := getg()  // 获取接收者协程
      // 开始封装sudog
      mysg := acquireSudog()
      mysg.releasetime = 0
      mysg.elem = ep  // 接收者的数据地址
      mysg.waitlink = nil
      gp.waiting = mysg
      mysg.g = gp
      mysg.selectdone = nil
      gp.param = nil
      c.recvq.enqueue(mysg)  // 放入recvq中
      
      // 调用goparkunlock，阻塞当前接收者协程。设置状态为waiting。
      goparkunlock(&c.lock, "chan receive", traceEvGoBlockRecv, 3)

      
      // 若能到此，表示接收者协程被某个发送者唤醒
      if mysg != gp.waiting {
         throw("G waiting list is corrupted!")
      }
      gp.waiting = nil
      haveData := gp.param != nil
      gp.param = nil
      releaseSudog(mysg)

      // 从发送者接收到数据。已经被写入到sudo.elem（即ep地址）了。
      if haveData {
         selected = true
         received = true
         return
      }

      lock(&c.lock)
      if c.closed == 0 {
         throw("chanrecv: spurious wakeup")
      }
      return recvclosed(c, ep)
   }

   // 如果是buffered channel
}
```

总结：

1. 如果是从closed channel中接收数据，则返回接收为false。
2. 尝试从`sendq`获取之前被阻塞的发送者。
3. 如果有发送者协程等待发送，则：
   1. 从发送者sudog中拷贝发送者的数据；
   2. 调用`goready`准备唤醒发送者协程。
4. 如果当前没有发送者协程准备好，则：
   1. 封装接收者协程为sudog，并放入`recvq`中；
   2. 调用`goparkunlock`将该协程置成waiting等待。

