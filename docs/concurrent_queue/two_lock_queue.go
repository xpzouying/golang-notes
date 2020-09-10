package main

import "sync"

type TwoLockQueue struct {
	Head *Node
	Tail *Node

	headlock sync.Mutex
	taillock sync.Mutex
}

func NewTwoLockQueue() *TwoLockQueue {
	node := &Node{}

	return &TwoLockQueue{Head: node, Tail: node}
}

func (q *TwoLockQueue) Enqueue(v int) {
	node := &Node{Value: v}

	q.taillock.Lock()
	q.Tail.Next = node
	q.Tail = node
	q.taillock.Unlock()
}

func (q *TwoLockQueue) Dequeue() (int, error) {
	q.headlock.Lock()

	// if empty queue
	if q.Head.Next == nil {
		q.headlock.Unlock()
		return 0, ErrEmptyQueue
	}

	v := q.Head.Next.Value
	q.Head = q.Head.Next
	q.headlock.Unlock()
	return v, nil
}
