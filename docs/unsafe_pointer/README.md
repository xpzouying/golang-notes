# unsafe.Pointer 

## 通用访问Tips

```go
// equivalent to f := unsafe.Pointer(&s.f)
f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f))

// equivalent to e := unsafe.Pointer(&x[i])
e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0]))
```


## `unsafe.Pointer`访问数组

当前有个`[]int`数组对象，如何通过`unsafe.Pointer`直接访问对象中的元素。



## `unsafe.Pointer`访问struct的元素

可以通过偏移访问struct的元素。

```go
p = unsafe.Pointer(uintptr(p) + offset)
```

