###
在学云原生课程之前，我对kubernetes的了解仅仅停留在最基础的pod, deployment 等的使用，而对其它的分层架构，组件以及API设计原则知之甚少。在学习课程之后，我收获了：

---

- **[基础类]** Golang的常用语法，线程加锁调度以及内存管理
- **[基础类]** kubernates 的主要组件有哪些？
    - etcd如何实现key-value 存储， API Server 如何减少集群对etcd 的访问
    - 控制器Control Manager 如何协同工作，以及如何确保状态的一致性
    - KubeScheduler的职责是什么，怎么为待调试的pod选择合适的计算节点
    - Kubelet 的作用是什么以及Kube proxy基于不同插件的负载平稳实现有哪些？
- **[基础类]** Kubernetes 核心对象: Node , Pod, Namespace, ConfigureMap, Secret, Service, Replica set, Stateful set etc

---
- **[课程实践类]**
    - 通过Kube Admin搭建K8S集群；
    - Docker容器分段构建；
    - 编写Golang http server 并发布到K8S集群（优雅启动和优雅终止);
    - 使用Service, Ingress 发布到K8S集群;
    - 集成Promethus,使用Grafana dashboard 展现metric
    - Helm 管理器K8S spec;

- **[真实项目实践类]**

    - 使用K8S job/cronjob 解耦原有项目任务调度，增加job级别资源控制；
    - 迁移Spark hadoop集群到k8s, 进行性能对比测试；
    - 使用Helm 管理项目spec
    