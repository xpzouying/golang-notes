# gevent vs goroutine

最近在接触python2.7版本中的gevent多线程，所以就想记录gevent多线程模型的性能和Golang goroutine性能之间的区别。

## python代码

使用gevent运行。

```python
from gevent import monkey; monkey.patch_all()
import gevent

def f(i):
    sum = 0
    for i in range(1000000000):
        x = i
        x += 1
        x *= 2
        x /= 3
        x -= 5

        sum += x

    print sum


gevent.joinall([
        gevent.spawn(f, 1),
        gevent.spawn(f, 2),
        gevent.spawn(f, 3),
])
```


```bash
# time python ./gevent_demo.py
333328333333
333328333333
333328333333

real    0m0.592s
user    0m0.450s
sys     0m0.040s
```

## go代码

```go
package main

import (
	"log"
	"sync"
)

func f() {
	sum := 0
	for i := 0; i < 1000000000; i++ {
		x := i
		x += 1
		x *= 2
		x /= 3
		x -= 5

		sum += x
	}

	log.Printf("sum=%d", sum)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			f()
			wg.Done()
		}()
	}

	wg.Wait()
}
```


运行

```bash
# go build ./goroutine_demo.go

# time ./goroutine_demo
2020/10/27 16:36:32 sum=333333328333333333
2020/10/27 16:36:32 sum=333333328333333333
2020/10/27 16:36:32 sum=333333328333333333

real    0m1.036s
user    0m3.082s
sys     0m0.008s
```