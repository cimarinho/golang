package golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func Send(teste Teste) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.1.38"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "teste_go_topic"

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(teste)

	 // this is the []b
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          reqBodyBytes.Bytes(),
		}, nil)


	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

type Teste struct{
	Nome string
	Idade int8
}