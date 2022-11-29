package kafkaconsumer

// import (
// 	"exam/review_service/config"
// 	"exam/review_service/kafka_consumer/server"
// 	"fmt"
// )

// type KafkaCon struct {
// 	KafkaConn *server.KafkaConnConsumer
// }

// type KafkaConI interface {
// 	Reads() *server.KafkaConnConsumer
// }

// func NewKafkaReader(cfg config.Config) (KafkaConI, func(), error) {
// 	kafReader, err := server.NewKafkaConsumer(cfg)
// 	if err != nil {
// 		fmt.Println("error while kafReader, err := client.NewKafkaConsumer(cfg): ", err)
// 		return nil, nil, err
// 	}

// 	return &KafkaCon{
// 			KafkaConn: kafReader,
// 		}, func() {
// 			kafReader.Reader.Close()
// 		}, nil

// }

// func (k *KafkaCon) Reads() *server.KafkaConnConsumer {
// 	return k.KafkaConn
// }
