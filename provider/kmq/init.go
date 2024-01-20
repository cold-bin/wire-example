package kmq

import (
	"github.com/IBM/sarama"
	"sync"
)

var (
	kmqSyncProducer sarama.SyncProducer
	once            = &sync.Once{}
)

func KmqClient() sarama.SyncProducer {
	if kmqSyncProducer != sarama.SyncProducer(nil) {
		return kmqSyncProducer
	}

	once.Do(func() {
		config := sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
		config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
		config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

		// 构造一个消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = "web_log"
		msg.Value = sarama.StringEncoder("this is a test log")
		// 连接kafka
		client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
		if err != nil {
			panic(err)
		}
		kmqSyncProducer = client
	})

	return kmqSyncProducer
}
