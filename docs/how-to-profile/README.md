# Golang调优指南

常常需要对Golang进行优化，在这里总结一下Golang常用的调优技巧和流程。

## 调优步骤

1. 生成profile数据文件
2. 使用工具分析profile数据文件
3. 优化代码
4. 重复步骤1


## 详细过程

### 如何生成profile数据文件

常用的生成数据文件的方式有三种：

**1、挂载profile handler到HTTP Server上。**
  该方法常常应用到HTTP Server上，监听**运行中**的服务状态。基本步骤为：
  1. 挂载net/http/pprof到HTTP Handler
  2. 启动服务
  3. 请求`debug/pprof`相应的接口输出数据文件

  详细内容参考：[net/http/pprof](https://golang.org/pkg/net/http/pprof/) 

**2、使用`runtime/pprof`输出数据文件。**
  该方法的基本原理是，找到需要调优的热点代码块，
  1. 在进入代码块时，开始记录；
  2. 在结束代码块时，结束记录。

  详细内容参考：[runtime/pprof](https://golang.org/pkg/runtime/pprof/)

**3、编写压力测试用例（Benchmark）输出数据文件。**
  该方法的基本原理是：对热点函数编写对应的压测用例，运行压测输出profile数据文件。基本流程为：
  1. 编写benchmark测试用例
  2. 运行benchmark压测并生成profile数据文件：`go test -cpuprofile cpu.prof -memprofile mem.prof -bench .`

  详细内容参考：
  - [Golang testing/benchmark](https://golang.org/pkg/testing/#hdr-Benchmarks)
  - [强烈推荐 Brad Fizpatrick - Profiling & Optimizing in Go](https://github.com/bradfitz/talk-yapc-asia-2015/blob/master/talk.md)


### 如何分析profile数据文件

经过上面的流程，可以获取对应的profile数据文件。使用使用`go tool pprof cpu.prof/mem.prof`打开对应的数据文件。

工具界面如下，

![image](https://user-images.githubusercontent.com/3946563/94361230-67910480-00e5-11eb-98f3-46f926f0aea1.png)

常用的命令有：

- top：查看每个函数的耗时。
- top -cum：查看每个函数的累计耗时，包括函数内部的其他调用耗时。如果函数A中包括的函数B，那么函数A的耗时也包含了函数B的耗时。
- list：list命令后面接函数名，如`list func1`，查看函数func1中的详细消耗情况。
- web：打开浏览器，查看图形调用关系耗时。
- weblist：类似于list，浏览器中显示。

更详细的可以参考：
- [https://blog.golang.org/pprof](https://blog.golang.org/pprof)