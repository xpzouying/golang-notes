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
