package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"logCollect/logAgent/kafka"
	"logCollect/logAgent/tailf"
)

func main() {

	fmt.Println("开始")
	// 读取初始化配置文件
	filename := "E:\\Go\\logCollect\\logAgent\\conf\\logAgent.conf"
	err := loadInitConf("ini", filename)
	if err != nil {
		fmt.Printf("导入配置文件错误:%v\n", err)
		panic("导入配置文件错误")
		return
	}

	// 初始化日志信息
	err = initLogger()
	if err != nil {
		fmt.Printf("导入日志文件错误:%v\n", err)
		panic("导入日志文件错误")
		return
	}
	// 输出成功信息
	logs.Debug("导入日志成功%v", logConfig)

	// 初识化etcd
	collectConf, err := initEtcd(logConfig.etcdAddr, logConfig.etcdKey)
	if err != nil {
		logs.Error("初始化etcd失败", err)
	}
	logs.Debug("初始化etcd成功!")

	// 初始化tailf
	err = tailf.InitTail(collectConf, logConfig.chanSize)
	if err != nil {
		logs.Error("初始化tailf失败:", err)
		return
	}
	logs.Debug("初始化tailf成功!")

	// 初始化Kafka
	err = kafka.InitKafka(logConfig.KafkaAddr)
	if err != nil {
		logs.Error("初识化Kafka producer失败:", err)
		return
	}
	logs.Debug("初始化Kafka成功!")

	// 运行
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed:", err)
	}
	logs.Info("程序退出")
}
