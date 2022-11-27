package kafkaproducer

import (
	"exam/post_service/config"
	"exam/post_service/kafka_producer/client"
	"fmt"
)

type Kafka struct {
	KafkaFuncs *client.KafkaConn
}

type KafkaI interface {
	Funcs() *client.KafkaConn
}

func NewKafka(cfg config.Config) (KafkaI, func(), error) {
	kaf, err := client.NewKafkaConnection(cfg)
	if err != nil {
		fmt.Println("error while kaf, err := client.NewKafkaConnection(cfg): ", err)
		return nil, nil, err
	}

	return &Kafka{
			KafkaFuncs: kaf,
		}, func() {
			kaf.Conn.Close()
		}, nil
}

func (k *Kafka) Funcs() *client.KafkaConn {
	return k.KafkaFuncs
}
