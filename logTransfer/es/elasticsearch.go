package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type LogMessage struct {
	App     string
	Topic   string
	Message string
}

var (
	esClient *elastic.Client
)

func InitEs(addr string) (err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		fmt.Println("connect es error", err)
		return nil
	}
	esClient = client
	return
}

func SendToES(topic string, data []byte) (err error) {
	msg := &LogMessage{}
	msg.Topic = topic
	msg.Message = string(data)

	_, err = esClient.Index().
		Index(topic).
		BodyJson(msg).
		Do(context.Background())
	if err != nil {
		panic(err)
		return
	}
	return
}
