package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func initLogger() (err error) {

	config := make(map[string]interface{})
	config["filename"] = logConfig.logPath
	config["level"] = convertLogLevel(logConfig.logLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("初始化日志, 序列化失败:", err)
		return
	}
	_ = logs.SetLogger(logs.AdapterFile, string(configStr))

	return
}
