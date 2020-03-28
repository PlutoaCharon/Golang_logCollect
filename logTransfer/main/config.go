package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

type LogConfig struct {
	KafkaAddr  string
	KafkaTopic string
	EsAddr     string
	LogPath    string
	LogLevel   string
}

var (
	logConfig *LogConfig
)

func InitConfig(confType string, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Printf("初始化配置文件出错:%v\n", err)
		return
	}
	// 导入配置信息
	logConfig = &LogConfig{}
	// 日志级别
	logConfig.LogLevel = conf.String("logs::log_level")
	if len(logConfig.LogLevel) == 0 {
		logConfig.LogLevel = "debug"
	}
	// 日志输出路径
	logConfig.LogPath = conf.String("logs::log_path")
	if len(logConfig.LogPath) == 0 {
		logConfig.LogPath = "E:\\Go\\logCollect\\logTransfer\\logs\\my.log"
	}

	// Kafka
	logConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(logConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("初识化Kafka addr失败")
		return
	}
	logConfig.KafkaTopic = conf.String("kafka::topic")
	if len(logConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("初识化Kafka topic失败")
		return
	}

	// Es
	logConfig.EsAddr = conf.String("elasticsearch::addr")
	if len(logConfig.EsAddr) == 0 {
		err = fmt.Errorf("初识化Es addr失败")
		return
	}
	return
}
