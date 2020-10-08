# LRU Cache

LRU Cache的Golang实现。


为了平衡Cache查找和淘汰机制的性能，引入2个数据结构：

1. Hash Map：
   1. Key：表示插入元素的key，在题目中为int类型。
   2. Value：表示插入元素的地址，地址中保存Node对象。
2. 双向链表：
   1. 最近最久未使用的元素放在队首，最近使用的元素，包括get、put的元素更新到队尾。


**总体思路为：**

对于LRU Cache的操作有3个：

1. 初始化工作
   1. 初始化空双链表
   2. 初始化Hash Map
2. get操作：
   1. 若在cache中，则对应的值；更新该key到链表的队尾，表示最近访问过；
   2. 若不在cache中，则返回-1。
   3. 判断是否在lru cache中，可以直接通过hashmap以O(1)的时间复杂度定位。
   4. 更新节点操作的时间复杂度：由于我们hashmap value中，保存的是Node的地址，所以可以直接通过Node的Prev和Next更新链表位置。
3. put操作：
   1. 若key为现有的key，则更新value值，并更新Node的位置到链表的尾部。
   2. 若key为新的key：
      1. 若还未到达容量上限，则直接添加到链表尾部。
      2. 若达到容量上限，则淘汰链表头部节点，并把新节点插入到链表尾部。


**具体代码如下：**

```go
// Node 是双向链表节点
type Node struct {
    Key int
    Value int
    Prev  *Node
    Next  *Node
}

type LRUCache struct {
    Head *Node
    Tail *Node

    // Cap 为cache大小
    Cap   int

    // M: <node_value: node_address>
    M map[int]*Node

    mu sync.Mutex
}


func Constructor(capacity int) LRUCache {
    // 初始化链表状态
    head := new(Node)
    tail := new(Node)
    head.Next = tail
    tail.Prev = head

    return LRUCache{
        Head: head,
        Tail: tail,
        Cap:  capacity,
        M:    make(map[int]*Node, capacity),
    }
}


// Get 获取key对应的value。如果key不存在，则返回-1
func (c *LRUCache) Get(key int) int {
    c.mu.Lock()
    defer c.mu.Unlock()

    node, ok := c.M[key]
    if !ok { // 如果当前cache不存在
        return -1
    }

    if node.Next == c.Tail { // 如果已经是最后一个元素，则直接返回
        return node.Value
    }

    // 更新排序
    // 1. 把node节点从链表中移除
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev

    // 2. 添加到队尾
    c.Tail.Prev.Next = node
    node.Prev = c.Tail.Prev
    c.Tail.Prev = node
    node.Next = c.Tail

    return node.Value
}

func (c *LRUCache) Put(key int, value int) {
    c.mu.Lock()
    defer c.mu.Unlock()

    // 如果key已经存在，则直接返回
    if node, ok := c.M[key]; ok {
        if node.Value != value {
            node.Value = value
        }
        
        // 更新排序
        node.Prev.Next = node.Next
        node.Next.Prev = node.Prev

        c.Tail.Prev.Next = node
        node.Prev = c.Tail.Prev
        c.Tail.Prev = node
        node.Next = c.Tail

        return
    }

    node := &Node{Key: key, Value: value}

    // 如果没有到达上限，直接插入在队尾位置
    if len(c.M) < c.Cap {
        tailPrev := c.Tail.Prev

        tailPrev.Next = node
        node.Prev = tailPrev

        c.Tail.Prev = node
        node.Next = c.Tail

        c.M[key] = node // 保存在hash中

        return
    }

    // 如果已经达到上限，则进行淘汰。淘汰链表头部节点。

    // 从cache中需要删除队首的key
    old := c.Head.Next   // 需要被淘汰的节点
    delete(c.M, old.Key) // 先删除hash中的元素

    old.Key = key
    old.Value = value

    // 从链表中删除old节点
    c.Head.Next = old.Next
    old.Next.Prev = c.Head

    // 在此，为了复用元素内存，不创建新的元素。
    // 更新值
    // 插入到链表尾部
    c.Tail.Prev.Next = old
    old.Prev = c.Tail.Prev
    old.Next = c.Tail
    c.Tail.Prev = old

    // 更新hash
    c.M[key] = old
}
```


**运行结果：**


```
执行结果： 通过
执行用时： 120 ms , 在所有 Go 提交中击败了 92.89% 的用户
内存消耗： 11.6 MB , 在所有 Go 提交中击败了 94.39% 的用户
```