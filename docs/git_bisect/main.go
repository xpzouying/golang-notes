package main

import (
	"errors"
	"log"
)

func Mod(value, m int) (int, error) {

	if m == 0 {
		return 0, errors.New("divide zero")
	}

	return int(value % m), nil
}

func main() {
	got, err := Mod(5, 3)
	if err != nil {
		log.Println("Mod failed, divide zeror")
	}

	log.Println(got)
}
