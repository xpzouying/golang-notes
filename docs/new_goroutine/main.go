package main

func sum(a, b int) int {
	return a + b
}

func main() {
	v1 := 1
	v2 := 2

	go sum(v1, v2)
}

// GOOS=linux GOARCH=amd64 go build -gcflags "-N -l -S" main.go
