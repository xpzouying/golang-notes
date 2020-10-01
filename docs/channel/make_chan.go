package main

func main() {
	c1 := make(chan int, 10)
	c2 := make(chan int)
	_ = c1
	_ = c2
}
