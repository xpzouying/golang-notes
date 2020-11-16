package main

import "log"

func makeSlice() {
	l := make([]int, 10)
	for i := 0; i < 5; i++ {
		l = append(l, i)
	}

	log.Printf("%v", l)
}

func main() {
	makeSlice()
}
