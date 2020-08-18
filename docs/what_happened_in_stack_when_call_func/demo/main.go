package main

//go:noinline
func sum(a, b int) int {
	return a + b
}

func main() {
	a, b := 1, 2
	sum(a, b)
}
