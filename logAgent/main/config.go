package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"logCollect/logAgent/tailf"
)

var (
	logConfig *Config
)

// 日志配置
type Config struct {
	logLevel string
	logPath  string

	chanSize    int
	KafkaAddr   string
	collectConf []tailf.CollectConf

	etcdAddr string
	etcdKey  string
}

//// 日志收集配置
//func loadCollectConf(conf config.Configer) (err error) {
//	var c tailf.CollectConf
//
//	c.LogPath = conf.String("collect::log_path")
//	if len(c.LogPath) == 0 {
//		err = errors.New("无效的 collect::log_path ")
//		return
//	}
//
//	c.Topic = conf.String("collect::topic")
//	if len(c.Topic) == 0 {
//		err = errors.New("无效的 collect::topic ")
//		return
//	}
//
//	logConfig.collectConf = append(logConfig.collectConf, c)
//	return
//}

// 导入初始化配置
func loadInitConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Printf("初始化配置文件出错:%v\n", err)
		return
	}
	// 导入配置信息
	logConfig = &Config{}
	// 日志级别
	logConfig.logLevel = conf.String("logs::log_level")
	if len(logConfig.logLevel) == 0 {
		logConfig.logLevel = "debug"
	}
	// 日志输出路径
	logConfig.logPath = conf.String("logs::log_path")
	if len(logConfig.logPath) == 0 {
		logConfig.logPath = "E:\\Go\\logagent\\logs\\my.log"
	}

	// 管道大小
	logConfig.chanSize, err = conf.Int("logs::chan_size")
	if err != nil {
		logConfig.chanSize = 100
	}

	// Kafka
	logConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(logConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("初识化Kafka失败")
		return
	}

	// etcd
	logConfig.etcdAddr = conf.String("etcd::addr")
	if len(logConfig.etcdAddr) == 0 {
		err = fmt.Errorf("初识化etcd addr失败")
		return
	}

	logConfig.etcdKey = conf.String("etcd::configKey")
	if len(logConfig.etcdKey) == 0 {
		err = fmt.Errorf("初识化etcd configKey失败")
		return
	}

	//err = loadCollectConf(conf)
	//if err != nil {
	//	fmt.Printf("导入日志收集配置错误:%v", err)
	//	return
	//}
	return
}
