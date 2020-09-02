package main

import "unsafe"

func accessSlice() {
	// init slice
	a1 := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		a1[i] = 100 + i
	}

	// access
	for i := 0; i < 10; i++ {
		v := *(*int)(unsafe.Pointer(&a1[i]))

		println(v)
	}
}

func main() {
	accessSlice()
}
