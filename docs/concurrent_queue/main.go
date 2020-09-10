package main

import (
	"fmt"
	"log"
)

func demoSimpleQueue() {
	q := NewSimpleQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	i := 1
	for {
		v, err := q.Dequeue()
		if err != nil {
			log.Printf("dequeue finished: %v", err)
			break
		}

		if v != i {
			panic(fmt.Sprintf("dequeue value=%d, want=%d", v, i))
		}

		i++
	}
}

func main() {
	demoSimpleQueue()
}
