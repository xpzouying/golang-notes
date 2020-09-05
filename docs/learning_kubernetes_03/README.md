# Kubernetes 编排

## 为什么需要Pod？

当我们启动多个程序时，由于该程序之间有依赖关系，所以我们希望在集群中调度到同一台宿主机上，所以就引入了Pod来解决这种成组调度。

Pod时Kubernetes中原子调度单位，而不是资源。

容器之间的紧密协作称为：“超亲密关系”。典型特征包括但不限于：

- 相互之间的直接的文件交换
- 使用localhost或者Socket文件进行本地通信
- 非常频繁的远程调用
- 共享Linux Namespace
- 等等


## Pod 实现原理

在Pod中，Infra容器永远是第一个被创建的容器，然后再是用户容器通过`Join Network Namespace`的方式与其关联在一起。


对于一个Pod来说，由于Infra容器的存在，Infra容器维持基本的基础设施，也即是：

1. Pod内的容器：
   - 可以直接使用localhost通信
   - 共享Pod的网络设备
2. 对Pod自身来说：
   - 只有1个IP地址，即该Pod的Network Namespace对应的IP地址
   - 生命周期只跟Infra容器一致，与内部的其他容器实例无关


`sidecar`模式：
可以在一个Pod中，启动一个辅助容器，来完成一些独立于主进程（主容器）之外的工作。


**总结**

1. 容器和虚拟机

    容器和虚拟机完全不同。
    容器本质上来说是一个**进程**。
    虚拟机是模拟物理机环境。


2. Kubernetes的调度
   1. Pod是调度的基本单位。
   2. Pod中有个特殊的容器：`Init Container`，负责Pod的基础设施。

    Swarm项目只解决了单容器调度的问题，没有解决应用架构编排的问题。
    从而有了Kubernetes的出现。


3. sidecar模式
   可以启动一个辅助的容器，来完成独立于主容器之外的工作。


-----

## Kubernetes 作业管理


**Pod Status：状态**

- Pending：Pod的YAML文件已经提交给Kubernetes，API对象已经创建并保存在etcd中。但是有可能不能被顺利创建。
- Running：已经调度成功，跟具体的节点绑定。
- Succeeded：Pod中的所有容器都已经完毕，并已经退出了。
- Failed：Pod里至少有个容器以非正常的状态退出（非0的返回码）。
- Unknown：异常状态，意味着Pod的状态不能被持续的被kubelet汇报给kube-apiserver。可能是Master和Kubelet之间的通信出了问题。

