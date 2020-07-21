package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {
	go writer()
	go reader()
	select {}
}

//阅读器
func reader() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "topic-A",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	r.SetOffset(19) //从19位开始读取

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Println(m.Offset, string(m.Key), string(m.Value))
	}

	defer r.Close()
}

//写入器
func writer() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "topic-A",
		Balancer: &kafka.LeastBytes{},
	})
	strs := kafka.Message{Key: []byte("1"), Value: []byte("hello world")}
	err := w.WriteMessages(context.Background(), strs)
	if err != nil {
		panic(err)
	}

	defer w.Close()
}
