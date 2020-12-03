# 如何使用`go tool trace`调优Goroutine


## Demo1 串形化计算

为了demo方便，使用CPU密集型的代码：计算一系列的和计算。

```go
func seq() {
	count := 0.

	for i := 0.; i < 20480.; i += 0.2 {
		for j := 20480.; j > 0; j -= 0.2 {
			count += (i*i - j*j + i*j)
		}
	}

	log.Printf("count=%v", count)
}
```


**生成trace数据**

```go
func main() {
	// defer profile.Start(profile.TraceProfile).Stop()
	seq()
}
```

**运行及查看时间**

```bash
go build .


# 与此同时，注意CPU使用情况
➜  trace (master) ✗ time ./trace
2020/12/03 01:16:18 count=1.0994686781389624e+18
./trace  13.86s user 0.00s system 99% cpu 13.866 total
```

**打开trace**

```go
defer profile.Start(profile.TraceProfile).Stop()
```


## Demo3：并发运行

```go
func seq3() {
	count := 0.
	mu := new(sync.Mutex)

	wg := new(sync.WaitGroup)
	limit := make(chan struct{}, 4)

	for i := 0.; i < 20480.; i += 0.2 {
		i := i

		limit <- struct{}{}
		wg.Add(1)

		go func() {
			subCount := 0.
			for j := 20480.; j > 0; j -= 0.2 {

				subCount += (i*i - j*j + i*j)
			}

			mu.Lock()
			count += subCount
			mu.Unlock()

			wg.Done()
			<-limit
		}()
	}
	wg.Wait()

	log.Printf("count=%v", count)
}
```


运行：

```bash
➜  trace (master) ✗ time ./trace
2020/12/03 01:17:00 count=1.0994686780039835e+18
./trace  13.90s user 0.01s system 396% cpu 3.510 total
```


## Dem4：消费者生产者模型

```go

func handleByWorkers() {

	// 用来接收i
	ch := make(chan float64)

	go func() {
		for i := 0.; i < 20480.; i += 0.2 {
			ch <- i
		}

		close(ch)
	}()

	// 计算总数的worker
	chSubCount := make(chan float64) // subCount的数据流
	chRes := make(chan float64)      // 结果的数据流
	go func() {
		// 总数
		count := 0.

		for subCount := range chSubCount {
			count += subCount
		}

		// 完成计数
		chRes <- count
	}()

	// N个消费者
	// N := runtime.NumCPU()
	N := 4
	wg := new(sync.WaitGroup)
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			for i := range ch {
				subCount := 0.
				for j := 20480.; j > 0; j -= 0.2 {
					subCount += (i*i - j*j + i*j)
				}
				chSubCount <- subCount // 每次的计算结果需要提供给计数的worker
			}

			wg.Done()
		}()
	}

	wg.Wait()
	close(chSubCount) // 表示子结果都计算完成
	count := <-chRes

	log.Printf("count=%v", count)
}
```


```bash
➜  trace (master) ✗ time ./trace
2020/12/03 01:39:12 count=1.0994686780039835e+18
./trace  15.21s user 0.14s system 321% cpu 4.778 total
```


**使用trace优化**

* channel 增加buffer


```bash
➜  trace (master) ✗ time ./trace
2020/12/03 01:48:26 count=1.0994686780039837e+18
./trace  13.72s user 0.01s system 393% cpu 3.486 total
```