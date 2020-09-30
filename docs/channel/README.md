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
