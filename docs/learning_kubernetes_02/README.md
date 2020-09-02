# Kubernetes 集群搭建与实践


## 为什么不能容器部署Kubernetes？

因为对于`kubelet`组件，无法做到容器化。

对于`kubelet`本身，需要操作宿主机上的资源，文件系统、网络等等。

所以，kubeadm选择了一种妥协方案：**把kubelet直接运行在宿主机上，然后使用其他容器部署其他的Kubernetes组件。**


-----

## 定义一个Kubernetes应用

1. 制作容器的镜像

2. 编写配置文件

   ```bash
   kubectl create -f 我的配置文件
   ```

   配置文件定义了一个`API Object`。
   

Pod就是Kubernetes的“应用”，而一个应用，就可以由多个容器组成。



## 总结

### 练习流程

制作应用，发布Docker镜像。

1. 本地使用Docker测试代码，制作镜像；
2. 选择合适的Kubernetes API对象，编写对应的YAML文件。
3. 在Kubernetes上部署YAML文件

注意：部署Kubernetes后，不推荐使用Docker命令行了。