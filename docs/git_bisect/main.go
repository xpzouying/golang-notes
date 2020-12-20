package main

import "log"

func Mod(value, m int) int {

	return int(value % m)
}

func main() {
	log.Println(Mod(5, 3))
}
