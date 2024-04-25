package kafka

import (
	"encoding/json"
	"github.com/ODawah/url-shortener/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

const (
	KafkaGroupId = "Requests-data-service"
)

var ConsumerClient *kafka.Consumer

func InitializeConsumer() error {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_SERVER"),
		"security.protocol": os.Getenv("KAFKA_PROTOCOL"),
		"sasl.mechanisms":   os.Getenv("KAFKA_SASL_MECHANISM"),
		"sasl.username":     os.Getenv("KAFKA_USERNAME"),
		"sasl.password":     os.Getenv("KAFKA_PASSWORD"),
		"group.id":          KafkaGroupId,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return err
	}
	ConsumerClient = c
	topic := KafkaTopic
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return err
	}
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var requestData models.RequestData
			err = json.Unmarshal(msg.Value, &requestData)
			if err != nil {
				log.Printf("Error decoding message: %v\n", err)
				continue
			}
			log.Printf("Received Request: %+v\n", requestData)
		} else {
			log.Printf("Error: %v\n", err)
		}
	}
}
