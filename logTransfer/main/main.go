package main

import (
	"github.com/astaxie/beego/logs"
	"logCollect/logTransfer/es"
)

func main() {
	// 初始化配置
	err := InitConfig("ini", "E:\\Go\\logCollect\\logTransfer\\config\\logTransfer.conf")
	if err != nil {
		panic(err)
		return
	}
	logs.Debug("初始化配置成功")

	//初始化日志模块
	err = initLogger(logConfig.LogPath, logConfig.LogLevel)
	if err != nil {
		panic(err)
		return
	}
	logs.Debug("初始化日志模块成功")

	// 初始化Kafka
	err = InitKafka(logConfig.KafkaAddr, logConfig.KafkaTopic)
	if err != nil {
		logs.Error("初始化Kafka失败, err:", err)
		return
	}
	logs.Debug("初始化Kafka成功")

	// 初始化Es
	err = es.InitEs(logConfig.EsAddr)
	if err != nil {
		logs.Error("初始化Elasticsearch失败, err:", err)
		return
	}
	logs.Debug("初始化Es成功")

	// 运行
	err = run()
	if err != nil {
		logs.Error("运行错误, err:", err)
		return
	}

	logs.Warn("logTransfer 退出")
}
