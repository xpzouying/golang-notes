package main

import "errors"

var ErrEmptyQueue = errors.New("empty queue")

type Node struct {
	Value int
	Next  *Node
}

type Queue interface {
	Enqueue(value int)
	Dequeue() (int, error)
}
