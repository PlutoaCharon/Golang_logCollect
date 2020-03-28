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

[Golang实战之海量日志收集系统（一）项目背景介绍](https://blog.csdn.net/qq_43442524/article/details/105023724)

[Golang实战之海量日志收集系统（二）收集应用程序日志到Kafka中](https://blog.csdn.net/qq_43442524/article/details/105024906)

[Golang实战之海量日志收集系统（三）简单版本logAgent的实现](https://blog.csdn.net/qq_43442524/article/details/105027853)

[Golang实战之海量日志收集系统（四）etcd介绍与使用etcd获取配置信息](https://blog.csdn.net/qq_43442524/article/details/105044853)

[Golang实战之海量日志收集系统（五）根据etcd配置项创建多个tailTask](https://blog.csdn.net/qq_43442524/article/details/105054482)

[Golang实战之海量日志收集系统（六）监视etcd配置项的变更](https://blog.csdn.net/qq_43442524/article/details/105072952)

## logTransfer

更新ing

## logBeegoWeb 

更新ing