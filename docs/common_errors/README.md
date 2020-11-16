# 总结Golang常见错误


## make slice


```go
	l := make([]int, 10)
	for i := 0; i < 5; i++ {
		l = append(l, i)
	}

	log.Printf("%v", l)
```

运行程序，输出

```
0 0 0 0 0 0 0 0 0 0 0 1 2 3 4
```
