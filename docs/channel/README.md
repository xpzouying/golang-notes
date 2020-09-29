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
     