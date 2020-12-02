package main

import (
	"log"
	"sync"
)

func seq() {
	count := 0.

	for i := 0.; i < 20480.; i += 0.2 {
		for j := 20480.; j > 0; j -= 0.2 {
			count += (i*i - j*j + i*j)
		}
	}

	log.Printf("count=%v", count)
}

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

func handleByWorkers() {

	// 用来接收i
	ch := make(chan float64, 100)

	go func() {
		for i := 0.; i < 20480.; i += 0.2 {
			ch <- i
		}

		close(ch)
	}()

	// 计算总数的worker
	chSubCount := make(chan float64, 100) // subCount的数据流
	chRes := make(chan float64)           // 结果的数据流
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

func main() {
	// defer profile.Start(profile.TraceProfile).Stop()

	// seq3()

	handleByWorkers()
}
