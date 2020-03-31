# 海量日志收集服务logCollect 

## 项目架构图
![项目架构图](https://pic.downk.cc/item/5e7f4609504f4bcb047fe140.jpg)

## 说明
logAgent从通过tailf从Etcd中获取要收集的日志信息从业务服务器读取日志信息，发往Kafka

logTransfer负责从Kafka读取日志，写入到Elasticsearch中

通过Kibana进行日志检索

最后通过Web界面控制Etcd管理日志配置

-------

## logAgent
![logAgent架构](https://pic.downk.cc/item/5e7ca7f0504f4bcb04c02c6b.jpg)

### logAgent主要实现的功能
- 可以自行配置要收集的日志文件

- 从Etcd中获取日志收集项

- 读取日志文件

- 写入到Kafka中

- logAgent可以同时运行多个日志收集任务

- 实现实时配置项变更

- 根据当前服务器的IP地址获取配置项

### 博客讲解

[Golang实战之海量日志收集系统（一）项目背景介绍](https://plutoacharon.github.io/2020/03/29/Golang%E5%AE%9E%E6%88%98%E4%B9%8B%E6%B5%B7%E9%87%8F%E6%97%A5%E5%BF%97%E6%94%B6%E9%9B%86%E7%B3%BB%E7%BB%9F%EF%BC%88%E4%B8%80%EF%BC%89%E9%A1%B9%E7%9B%AE%E8%83%8C%E6%99%AF%E4%BB%8B%E7%BB%8D/)

[Golang实战之海量日志收集系统（二）收集应用程序日志到Kafka中](https://plutoacharon.github.io/2020/03/29/Golang%E5%AE%9E%E6%88%98%E4%B9%8B%E6%B5%B7%E9%87%8F%E6%97%A5%E5%BF%97%E6%94%B6%E9%9B%86%E7%B3%BB%E7%BB%9F%EF%BC%88%E4%BA%8C%EF%BC%89%E6%94%B6%E9%9B%86%E5%BA%94%E7%94%A8%E7%A8%8B%E5%BA%8F%E6%97%A5%E5%BF%97%E5%88%B0Kafka%E4%B8%AD/)

[Golang实战之海量日志收集系统（三）简单版本logAgent的实现](https://plutoacharon.github.io/2020/03/29/Golang%E5%AE%9E%E6%88%98%E4%B9%8B%E6%B5%B7%E9%87%8F%E6%97%A5%E5%BF%97%E6%94%B6%E9%9B%86%E7%B3%BB%E7%BB%9F%EF%BC%88%E4%B8%89%EF%BC%89%E7%AE%80%E5%8D%95%E7%89%88%E6%9C%AClogAgent%E7%9A%84%E5%AE%9E%E7%8E%B0/)

[Golang实战之海量日志收集系统（四）etcd介绍与使用etcd获取配置信息](https://plutoacharon.github.io/2020/03/29/Golang%E5%AE%9E%E6%88%98%E4%B9%8B%E6%B5%B7%E9%87%8F%E6%97%A5%E5%BF%97%E6%94%B6%E9%9B%86%E7%B3%BB%E7%BB%9F%EF%BC%88%E5%9B%9B%EF%BC%89etcd%E4%BB%8B%E7%BB%8D%E4%B8%8E%E4%BD%BF%E7%94%A8etcd%E8%8E%B7%E5%8F%96%E9%85%8D%E7%BD%AE%E4%BF%A1%E6%81%AF/)

[Golang实战之海量日志收集系统（五）根据etcd配置项创建多个tailTask](https://plutoacharon.github.io/2020/03/29/Golang%E5%AE%9E%E6%88%98%E4%B9%8B%E6%B5%B7%E9%87%8F%E6%97%A5%E5%BF%97%E6%94%B6%E9%9B%86%E7%B3%BB%E7%BB%9F%EF%BC%88%E4%BA%94%EF%BC%89%E6%A0%B9%E6%8D%AEetcd%E9%85%8D%E7%BD%AE%E9%A1%B9%E5%88%9B%E5%BB%BA%E5%A4%9A%E4%B8%AAtailTask/)

[Golang实战之海量日志收集系统（六）监视etcd配置项的变更](https://plutoacharon.github.io/2020/03/29/Golang%E5%AE%9E%E6%88%98%E4%B9%8B%E6%B5%B7%E9%87%8F%E6%97%A5%E5%BF%97%E6%94%B6%E9%9B%86%E7%B3%BB%E7%BB%9F%EF%BC%88%E5%85%AD%EF%BC%89%E7%9B%91%E8%A7%86etcd%E9%85%8D%E7%BD%AE%E9%A1%B9%E7%9A%84%E5%8F%98%E6%9B%B4/)

## logTransfer

项目架构图:

![logTransfer架构](https://pic.downk.cc/item/5e802310504f4bcb040062cd.jpg)

项目逻辑图: 

![logTransfer架构](https://pic.downk.cc/item/5e8022d8504f4bcb04003b10.jpg)

### logTransfer主要实现的功能

- 将日志数据写入到Kafka中

- 将消费的数据落地到Elastciseartch中

- 通过Kibana进行展示


### 博客讲解

[Golang实战之海量日志收集系统（七）logTransfer之从kafka中获取日志信息 ](https://plutoacharon.github.io/2020/03/29/Golang%E5%AE%9E%E6%88%98%E4%B9%8B%E6%B5%B7%E9%87%8F%E6%97%A5%E5%BF%97%E6%94%B6%E9%9B%86%E7%B3%BB%E7%BB%9F%EF%BC%88%E4%B8%83%EF%BC%89logTransfer%E4%B9%8B%E4%BB%8Ekafka%E4%B8%AD%E8%8E%B7%E5%8F%96%E6%97%A5%E5%BF%97%E4%BF%A1%E6%81%AF/)

[Golang实战之海量日志收集系统（八）logTransfer之将日志入库到Elasticsearch并通过Kibana进行展示 ](https://plutoacharon.github.io/2020/03/29/Golang%E5%AE%9E%E6%88%98%E4%B9%8B%E6%B5%B7%E9%87%8F%E6%97%A5%E5%BF%97%E6%94%B6%E9%9B%86%E7%B3%BB%E7%BB%9F%EF%BC%88%E5%85%AB%EF%BC%89logTransfer%E4%B9%8B%E5%B0%86%E6%97%A5%E5%BF%97%E5%85%A5%E5%BA%93%E5%88%B0Elasticsearch%E5%B9%B6%E9%80%9A%E8%BF%87Kibana%E8%BF%9B%E8%A1%8C%E5%B1%95%E7%A4%BA/)

## logBeegoWeb 

http://localhost:8080/index

项目管理:

![](https://pic.downk.cc/item/5e82f87e504f4bcb04068fc6.jpg)

项目申请:

![](https://pic.downk.cc/item/5e82f8e1504f4bcb0406f543.jpg)


日志列表:

![](https://pic.downk.cc/item/5e82f903504f4bcb04070eac.jpg)


日志申请:
![](https://pic.downk.cc/item/5e82f926504f4bcb04072767.jpg)
























