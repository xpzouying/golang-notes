# 深入剖析kubernetes - 学习笔记

本文主要是记录我学习《极客时间》中，对于张磊的《深入剖析kubernetes》这门课程的学习笔记。

## 预习篇 - Docker

### 问题1：遇到什么问题，解决什么问题，如何解决问题？

在云计算中（PaaS），如何打包是一个棘手的问题。

PaaS平台虽然有部署脚本，但是没有解决应用服务对于不同的环境的适配，常常导致对于不同的环境需要：重复打包、重复配置等流程。

Docker的杀手锏特性——Docker镜像。

对于Docker来说，通过系统级别的打包，使用了Cgroups和Namespace进行了隔离，这样就可以使得应用服务的发布变成了批量生成的过程。

但是对于Docker来说没有解决编排、管理的问题，所以就产生了其他的工具，比如：Deis、Flynn、Swarm等等。

**Swarm项目的根本目标：如何让开发者把应用部署在我的项目上。**

### 问题2：如何解决交付

解决了打包问题，但是最终的目标还是把用户的服务使用起来，那么下一个问题就是如何把服务部署到用户的环境中，单机、集群的部署、管理。

- 单机Docker：
    ```
    docker run “我的容器”
    ```

- 多机Docker：
    ```
    docker run -H “swarm address API” “我的容器”
    ```
    
    1. 创建容器的请求会被Swarm拦截下来
    2. 通过具体的调度算法找到合适的Docker Daemon运行起来


Docker中的编排项目（Container Orchestration）：Fig，后改名为Compose。

**Docker运行时：**

- Docker项目的运行时部分：Containerd
- Moby社区维护版本。

**Kubernetes的策略：**

微服务治理：Istio
有状态应用部署框架：Operator


-----


## 白话容器基础

- Docker项目通过“容器镜像”，解决了应用打包这个根本性难题；
- **容器本身没有价值，有价值的是“容器编排”；**



示例：命令：

```bash
docker run -it busybox /binsh
```

其中：`-it参数`：启动容器后，在容器内执行/bin/sh，分配一个文本输入/输出环境，也即是TTY。表示可以和Docker容器进行交互。



**Namespace机制：** Docker在使用`clone()`创建新进程的一个可选参数：`CLONE_NEWPID`。

`int pid = clone(main_function, stack_size, SIGCHLD, NULL); `
变成了
`int pid = clone(main_function, stack_size, CLONE_NEWPID | SIGCHLD, NULL); `

`CLONE_NEWPID`表示：在新进程空间里，PID是1。但在宿主进程空间里，PID还是为真实的数值，比如为100。



**PID Namespace：**使得每个Namespace里的应用进程，都会认为自己是当前容器里的第1号进程。还有Mount、UTS、IPC、Network、User这些Namspace。
所以，容器只是一个特殊的进程而已。

![image-20200828090220847](image-20200828090220847.png)

左边是虚拟机的工作原理，右边是容器的工作原理。



在理解了 Namespace 的工作方式之后，你就会明白，跟真实存在的虚拟机不同，在使用 Docker 的时候，并没有一个真正的“Docker 容器”运行在宿主机里面。Docker 项目帮助用户启动的，还是原来的应用进程，只不过在创建这些进程时，Docker 为它们加上了各种各样的 Namespace 参数。

虚拟机和容器的比对：

- 容器的优点：敏捷、高性能

- 虚拟机的优点：隔离彻底



### 容器的隔离

1. 容器只是运行在宿主机上的一个特殊的进程，那么多个容器使用的还是同一个宿主机的操作系统内核。

   容器里通过Mount Namespace单独挂在其他不同版本的操作系统文件，并不能改变宿主机内核的事实。**所以，在Windows宿主机上面运行Linux容器，或者在低版本的Linux宿主机上面运行高版本的Linux容器，都是不行的。**

2. 很多资源和对象不能被Namespace化。最典型的就是：时间。

   比如**settimeofday(2)**系统调用修改时间，整个宿主机的时间就被修改。

   可以使**用Seccomp**等技术，对容器内部发起的系统调用进行过滤，但是问题是：

   	- 无法知道应该开启和关闭哪些系统调用
   	- 多一层系统调用的过滤甄别，一定会拖累容器的性能

   **所以在生产环境，没有人敢把运行在物理机上的Linux容器直接暴露在公网上。**



### 容器的“限制”

Linux Cgroups全称是：Linux Control Group。主要作用是：限制一个进程组能够使用的资源上限，包括CPU、内存、磁盘、网络带宽等等。

还可以对进程进行优先级设置、审计、以及进程挂起和恢复等操作。

每一个docker容器启动的时候，都会在/sys/fs/cgroup/cpu/docker/目录下创建一个容器命名的文件夹，该文件夹下面是对于该容器的“控制组”，配置对于该容器的资源限制文件。



**宿主机上Cgroups DEMO**

1. 运行命令，占满CPU
```bash
while : ; do : ; done &[1] 226
# 该进程PID为226。
```

2. 使用top查看当前CPU占用。由于没有限制，所以当前CPU占用100%

3. 使用Cgroups进行限制CPU：

    - 进入目录/sys/fs/cgroup/cpu/docker，查看当前CPU信息：

    ```
    cat cpu.cfs_quota_us
    # -1。表示，使用CPU配额没有限制。
    ```

    ```
    cat cpu.cfs_period_us
    # 100000：CPU周期定义为100ms
    ```
        
    - 进行修改CPU使用配额：
    ```
    echo 20000 > /sys/fs/cgroup/cpu/docker/cpu.cfs_quota_us
    # 每周期使用20000us（即20ms），也就是占用20%的CPU。
    ```

    - 设置指定的PID进行限制：
    ```bash
    echo 226 > /sys/fs/cgroup/cpu/docker/tasks
    # 之前启动的时候，PID为226。
    ```

    - 查看限制是否生效：`top`



**docker的Cgroups DEMO**

1. 查看CPU周期：
    ```
    cat /sys/fs/cgroup/cpu/cpu.cfs_period_us
    # 默认为100ms（100000us）
    ```

2. 创建一个容器示例：
```bash
docker run -it --cpu-period=100000 --cpu-quota=20000 ubuntu bash
```

    **说明：**
    - cpu-quota参数指定对于一个周期（由--cpu-period）中，多少时间可以使用CPU资源。所以对于上面的示例来说，每个CPU周期，可以用到20%的CPU时间。

    - 在目录/sys/fs/cgroup/cpu/docker中，创建了以docker实例id为文件夹的“控制组”

3. 在容器中，让该容器的CPU跑满：
```bash
while : ; do : ; done &
```

4. 在宿主机上，查看CPU使用情况：top。能看到CPU占用20%。


### 总结

一个正在运行的Docker容器，其实就是启用了多个`Linux Namespace`（CPU、IO等）的应用进程，而这个进程能够使用的资源量，则受到`Cgroups`配置的限制。

所以，**容器是一个“单进程”模型**。

由于容器的本质是一个进程，用户的应用进程实际上就是容器里的`PID=1`的进程，这也就是后续创建的所有进程的父进程。也就意味着，在一个容器中，没有办法运行两个不同的应用，除非能事先找到`PID=1`的程序作为不同应用的父进程。这也是为什么很多人都会使用`systemd`或者`supervisord`替代应用本身作为容器的启动进程。


**Cgroups对资源限制能力的不足**

`/proc`文件系统的问题：

- Linux中`/proc`目录记录当前内核运行状态的一系列特殊文件，查看系统及当前正在运行进程的信息，比如CPU使用情况、内存占用率等，这些也是`top`查看系统信息的主要数据来源：top从`/proc/stat`中获取数据。

- 在容器里执行top指令，显示的是宿主机的CPU和内存数据，而不是当前容器的数据。原因是：`/proc`文件系统并不知道用户通过Cgroups给当前容器做了什么样的资源限制。

> 上述的问题可以借助`lxcfs`解决。
> 由于`top`是从`/proc/stat`下获取数据，所以让容器不挂在宿主机的该目录即可。




### Docker镜像

前面已经提到，对于Docker实例来说，简单来说是一个应用进程。该进程通过Namespace进行各种**隔离**，通过Cgroups进行资源的**限制**。

有个比较特殊的Namespace是文件系统的隔离：使用`Mount Namespace`。需要手动执行触发。`Mount Namespace`是基于对`chroot`的不断改良发明出来的，也是Linux中第一个Namespace。


关于Namespace的详细介绍可以参考改文章：
[左耳耗子 - DOCKER基础技术：LINUX NAMESPACE（上）](https://coolshell.cn/articles/17010.html)


**“容器镜像”**

挂在在容器根目录上，用于为容器进程提供隔离后执行环境的文件系统。也叫做：**rootfs（根文件系统）**。

一般包括：`/bin`, `/etc`, `/proc`等等。


**Docker项目**

最核心的原理就是为待创建的用户进程：

1. 启用Linux Namespace配置；

2. 设置指定的Cgroups参数；

3. 切换进程的根目录（Change Root）；（在该步骤会优先调用pivot_root系统调用）


rootfs只包含文件、配置和目录，不包含操作系统内核。

同一台机器上，所有的容器都共享操作系统的内核。


**依赖库/镜像版本**

对于应用程序来说，操作系统本身才是最完整的“依赖库”。

有了`rootfs`后，我们就可以对应用程序打包一套完整的依赖库环境。解决了完整的依赖问题后，现在需要解决的是该依赖库的版本控制问题。

举例来说，但对于传统的依赖库环境来说，v1.1是基于v1.0版本修改而成，v1.0版本和v1.1版本都具备完整的环境，它们之间没有太多的联系。未来出现v1.2版本后，经过少量修改后，又得生成一套全新的依赖库。那么，解决方法就是希望每次只需要维护修改的增量内容，而不是每次重新创建一个全新、完整的依赖库。

为此，引入了`layer`。每一步操作，会生成一个层，也即是增量rootfs。为此引入“联合文件系统（Union File System）”，UnionFS。


`AuFS`: Another UnionFS，又改名为：Alternative UnionFS，最后改为：Advance UnionFS

- 是对Linux原生UnionFS的重写和改进；

- 只能在Ubuntu和Debian发行版上使用；


对于AuFS，最关键的目录结构是在：`/var/lib/docker`下的diff目录：

```
/var/lib/docker/aufs/diff/<layer_id>
```


“镜像”实际上就是操作系统的rootfs，内容就是操作系统的所有文件和目录。不同的是：往往由多个“层layer”组成。

使用命令可以查看，

```bash
docker image inspect ubuntu:latest

//...
        "RootFS": {
            "Type": "layers",
            "Layers": [
                "sha256:918b1e79e35865cfaa7af9e07fa2a7aaaa2885e6bee964691a93c5db631b0aff",
                "sha256:83b575865dd109e77301a1be1e510cfffa6b89b9ff6355df22b5008315778263",
                "sha256:153bd22a8e96919e8eb890cc50aba51d7c16ea0746c2f020f21312f88e65f5c8",
                "sha256:ca893d4b83a60ef4e859785bc6b4072242ae07c7d6d0a07098847bc281b525b8",
                ...
            ]
        },
// ...
```

