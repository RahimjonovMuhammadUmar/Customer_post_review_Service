package server

import (
	"context"
	"exam/review_service/config"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KafkaConnConsumer struct {
	Reader    *kafka.Reader
	ConnClose func()
	Cfg       config.Config
}

func NewKafkaConsumer(cfg config.Config) (*KafkaConnConsumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:29092", "kafka:9092"},
		Topic:     "ids",
		Partition: 0,
		MinBytes:  1e3,  //10KB
		MaxBytes:  10e6, //10MB
	})
	return &KafkaConnConsumer{
		Reader:    reader,
		ConnClose: func() { reader.Close() },
	}, nil
}

func (k *KafkaConnConsumer) ViewIds() error {
	for {
		m, err := k.Reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("error while reading message", err)
			return err
		}
		
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

	}
}
