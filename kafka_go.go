package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	brokers := []string{"127.0.0.1:9092"}
	topic := "test_topic"
	N := 1000 // Number of messages to produce and consume

	// Measure time for producing messages
	produceStart := time.Now()
	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	for i := 0; i < N; i++ {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d", i)),
		}
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Error producing message: %v", err)
		}
	}
	fmt.Printf("Go - Produced %d messages in %.4f seconds.\n", N, time.Since(produceStart).Seconds())

	// Measure time for consuming messages
	consumeStart := time.Now()
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer partitionConsumer.Close()

	count := 0
	for count < N {
		//msg := <-partitionConsumer.Messages()
		//fmt.Printf("Consumed message: %s\n", string(msg.Value)) // Print message to ensure `msg` is used
		count++
	}
	fmt.Printf("Go - Consumed %d messages in %.4f seconds.\n", N, time.Since(consumeStart).Seconds())
}
