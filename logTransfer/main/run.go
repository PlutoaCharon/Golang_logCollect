package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"logCollect/logTransfer/es"
	"time"
)

func run() (err error) {

	// Kafka消费数据
	partitionList, err := kafkaClient.Client.Partitions(kafkaClient.Topic)
	if err != nil {
		logs.Error("Failed to get the list of partitions: ", err)
		return
	}
	for partition := range partitionList {
		pc, errRet := kafkaClient.Client.ConsumePartition(kafkaClient.Topic, int32(partition), sarama.OffsetNewest)
		if errRet != nil {
			err = errRet
			logs.Error("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {

			for msg := range pc.Messages() {
				logs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				err = es.SendToES(kafkaClient.Topic, msg.Value)
				if err != nil {
					logs.Warn("send to es failed, err:%v", err)
				}
			}

		}(pc)
	}

	time.Sleep(time.Hour)
	return
}
