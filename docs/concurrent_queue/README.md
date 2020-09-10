# 并发安全的队列

最近看了一篇关于**并发安全队列**的论文，在此记录该笔记。

## 参考论文

- 《Simple, Fast, and Practical Non-Blocking and Blocking Concurrent Queue Algorithms》 by Maged M. Michael, Michael L. Scott


-----

首先，定义队列数据结构。其中包括两部分：

1. Node节点

```go
type Node struct {
	Value int
	Next  *Node
}
```

2. 队列interface

```go
type Queue interface {
	Enqueue(value int)
	Dequeue() (int, error)
}
```


## 第一版：链表


**定义**

定义最简单的FIFO的链表。

```go
type SimpleQueue struct {
	Head *Node
	Tail *Node
}
```

其中数据结构包含：

- Head：指向头节点
- Tail：指向队尾节点


**入队**

```go
func (q *SimpleQueue) Enqueue(v int) {
	node := &Node{Value: v}
	q.Tail.Next = node
	q.Tail = node
}
```

1. 在队尾添加节点
2. 更新Tail，指向新的节点


**出队**

```go
func (q *SimpleQueue) Dequeue() (int, error) {
	// if empty queue
	if q.Head.Next == nil {
		return 0, ErrEmptyQueue
	}

	newHead := q.Head.Next
    retVal := newHead.Value

	oldHead := q.Head
	oldHead.Next = nil
	oldHead = nil

	q.Head = newHead
	return retVal, nil
}
```

1. 如果当前链表为空，则返回空链表错误。
2. 否则，
   1. 记录头部节点的值；
   2. 更新头部；
   3. 返回头部节点的值；