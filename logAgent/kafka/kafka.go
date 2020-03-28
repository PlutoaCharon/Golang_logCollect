package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error) {

	// Kafka生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 新建一个生产者对象
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("初识化Kafka producer失败:", err)
		return
	}
	logs.Debug("初始化Kafka producer成功,地址为:", addr)
	return
}

func SendToKafka(data, topic string) (err error) {

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)

	if err != nil {
		logs.Error("发送信息失败, err:%v, data:%v, topic:%v", err, data, topic)
		return
	}

	logs.Debug("read success, pid:%v, offset:%v, topic:%v\n", pid, offset, topic)
	return
}
