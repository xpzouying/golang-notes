package main

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

// type AddrNode struct {
// 	Value int
// 	Next  unsafe.Pointer
// }

type FreelockQueue struct {
	Head unsafe.Pointer
	Tail unsafe.Pointer
}

func NewFreelockQueue() *FreelockQueue {
	node := new(Node)

	return &FreelockQueue{
		Head: unsafe.Pointer(node),
		Tail: unsafe.Pointer(node),
	}
}

func (q *FreelockQueue) Enqueue(value int) {
	n := &Node{Value: value}

	for {
		tail := load(&q.Tail)
		next := tail.Next

		if tail == load(&q.Tail) { // 如果开始取的tail和现在取得还是一致的话？即没有新的node插入
			if next == nil { // 如果tail还是队尾最后一个节点，没有新的节点加入。
				// 尝试将新的Node加入到列表最后

				if atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&tail.Next)), // addr, 指针的指针
					unsafe.Pointer(next),                          // old
					unsafe.Pointer(n),                             // new
				) {
					return
				}
			}
		}
	}
}

func load(addr *unsafe.Pointer) *Node {
	return (*Node)(atomic.LoadPointer(addr))
}

func (q *FreelockQueue) Dequeue() (int, error) {

	return 0, errors.New("not implement")
}
