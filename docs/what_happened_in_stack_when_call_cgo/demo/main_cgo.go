package main

/*
int sum(int a, int b) { return a+b; }
*/
import "C"

func main() {
	a, b := C.int(1), C.int(2)
	C.sum(a, b)
}
