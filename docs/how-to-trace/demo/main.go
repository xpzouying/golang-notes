package main

import (
	"log"

	"github.com/pkg/profile"
)

func seq() {
	count := 0.

	for i := 0.; i < 102400.; i += 0.2 {
		for j := 102400.; j > 0; j -= 0.2 {
			count += (i*i - j*j + i*j)
		}
	}

	log.Printf("count=%v", count)
}

func main() {
	defer profile.Start(profile.TraceProfile).Stop()

	seq()
}
