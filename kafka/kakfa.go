package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

const (
	KafkaTopic = "requests-v1-topic"
)

var ProducerClient *kafka.Producer

func InitializeProducer() error {
	var err error

	ProducerClient, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_SERVER"),
		"security.protocol": os.Getenv("KAFKA_PROTOCOL"),
		"sasl.mechanisms":   os.Getenv("KAFKA_SASL_MECHANISM"),
		"sasl.username":     os.Getenv("KAFKA_USERNAME"),
		"sasl.password":     os.Getenv("KAFKA_PASSWORD"),
		"acks":              os.Getenv("ACKS"),
	})
	if err != nil {
		log.Println("error Connecting Kafka", err)
		return err
	}
	return nil
}
