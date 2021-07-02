# string 类型的值可以修改吗

参考来源：

- [string 类型的值可以修改吗](https://mp.weixin.qq.com/s/fl8WMsnfOxV0aD9Nic-F4w)
- [golang-string 和 bytes 之间的 unsafe 转换](https://jaycechant.info/2019/golang-unsafe-cast-between-string-and-bytes/)


```go
// 修改字符串的错误示例
func main() {
 x := "text"
 x[0] = "T"  // error: cannot assign to x[0]
 fmt.Println(x)
}
 
 
// 修改示例
func main() {
 x := "text"
 xBytes := []byte(x)
 xBytes[0] = 'T' // 注意此时的 T 是 rune 类型
 x = string(xBytes)
 fmt.Println(x) // Text
}
```