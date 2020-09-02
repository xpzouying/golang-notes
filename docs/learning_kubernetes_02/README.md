# Kubernetes 集群搭建与实践


### 为什么不能容器部署Kubernetes？

因为对于`kubelet`组件，无法做到容器化。

对于`kubelet`本身，需要操作宿主机上的资源，文件系统、网络等等。

所以，kubeadm选择了一种妥协方案：**把kubelet直接运行在宿主机上，然后使用其他容器部署其他的Kubernetes组件。**

