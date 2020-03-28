package main

import (
	"github.com/astaxie/beego/logs"
	"logCollect/logAgent/kafka"
	"logCollect/logAgent/tailf"
	"time"
)

func serverRun() (err error) {

	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("Send to Kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}

}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	//fmt.Printf("读取 msg:%s, topic:%s\n", msg.Msg, msg.Topic)
	_ = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
