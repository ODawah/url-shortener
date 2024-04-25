package kafka

import (
	"encoding/json"
	"github.com/ODawah/url-shortener/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func ProduceMessage(key string, msg models.RequestData) error {
	value, err := json.Marshal(msg)
	if err != nil {
		log.Printf("error marshalling message :%s", err)
		return err
	}
	topic := KafkaTopic
	err = ProducerClient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
		Key:            []byte(key),
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
