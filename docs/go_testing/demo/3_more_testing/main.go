package main

import "errors"

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("invalid input")
	}

	return a / b, nil
}
