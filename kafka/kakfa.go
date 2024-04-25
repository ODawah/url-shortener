package persistence

import (
	"encoding/json"
	"github.com/ODawah/url-shortener/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	KafkaServer = "localhost:9092"
	KafkaTopic  = "Requests-topic"
)

var ProducerClient *kafka.Producer

func InitializeProducer() error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
	})
	if err != nil {
		return err
	}
	ProducerClient = p
	return nil
}

func ProduceMessage(msg models.RequestData) error {
	value, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	topic := KafkaTopic
	err = ProducerClient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}
