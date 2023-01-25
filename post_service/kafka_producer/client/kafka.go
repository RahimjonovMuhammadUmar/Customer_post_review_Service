package client

import (
	"context"
	"encoding/json"
	"exam/post_service/config"
	"exam/post_service/genproto/post"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KafkaConn struct {
	Conn      *kafka.Conn
	ConnClose func()
}

func NewKafkaConnection(cfg config.Config) (*KafkaConn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.KafkaHost+":"+cfg.KafkaPort, "Created_posts", 0)
	if err != nil {
		fmt.Println("error while DialLeader", err)
		return &KafkaConn{}, err
	}
	return &KafkaConn{
		Conn: conn,
		ConnClose: func() {
			conn.Close()
		},
	}, err
}

func (k *KafkaConn) SendPost(req *post.PostWithoutReview) error {
	byteReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println("error while marshaling req")
		return err
	}
	_, err = k.Conn.WriteMessages(kafka.Message{
		Key: []byte("Post"),
		Value: byteReq,
	})
	return err

}
