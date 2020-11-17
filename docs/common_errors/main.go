package main

import (
	"log"
)

func makeSlice() {
	l := make([]int, 10)
	for i := 0; i < 5; i++ {
		l = append(l, i)
	}

	log.Printf("%v", l)
}

// 看似一道简单的 Go 题，考点不少：一道字节跳动面试题解析
// https://mp.weixin.qq.com/s/gQcCzrP8Pvr9WKONOl0ZHQ
func forGoroutine() {
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
		}()
	}

	log.Printf("total:%d sum %d", total, sum)
}

func main() {
	// makeSlice()

	forGoroutine()
}
