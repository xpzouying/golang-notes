# zyraft

Yet another raft implemention.

## 学习笔记

- [https://github.com/eliben/raft](https://github.com/eliben/raft)

这里有个系列文章讲述`raft`的实现。


## 思考题

读源码时，主要要考虑下列问题：

1. Leader的选举；

1. 日志的同步；

1. 日志的持久化；

1. 单元测试是如何编写的；