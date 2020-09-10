package main

type SimpleQueue struct {
	Head *Node
	Tail *Node
}

func NewSimpleQueue() *SimpleQueue {
	node := &Node{}

	return &SimpleQueue{Head: node, Tail: node}
}

func (q *SimpleQueue) Enqueue(v int) {
	node := &Node{Value: v}
	q.Tail.Next = node
	q.Tail = node
}

func (q *SimpleQueue) Dequeue() (int, error) {
	// if empty queue
	if q.Head.Next == nil {
		return 0, ErrEmptyQueue
	}

	newHead := q.Head.Next
	retVal := newHead.Value
	newHead.Value = 0

	oldHead := q.Head
	oldHead.Next = nil
	oldHead = nil

	q.Head = newHead
	return retVal, nil
}
