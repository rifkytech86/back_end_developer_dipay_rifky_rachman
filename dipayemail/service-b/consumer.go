package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"os/signal"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	//broker := "localhost:9092"
	//topic := "email-topic"

	fmt.Println("starting consumer")
	fmt.Println("starting consumer", broker, topic)
	fmt.Println("starting consumer")
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "email-consumer-group", // Change this to a unique group ID for your application
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	for {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			return
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				subject := string(e.Key)
				body := string(e.Value)

				err := sendEmail(subject, body)
				if err != nil {
					log.Printf("Error sending email: %v\n", err)
				} else {
					log.Printf("Email sent successfully. Subject: %s\n", subject)
				}

			case kafka.Error:
				log.Printf("Error: %v\n", e)

			default:
				log.Printf("Ignored event: %v\n", e)
			}
		}
	}
}

func sendEmail(subject, body string) error {

	fmt.Println("-------------------")
	fmt.Println(subject, body)
	fmt.Println("-------------------")
	// Configure your email sending here
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@example.com")
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.your-email-provider.com", 587, "your-email@example.com", "your-email-password")

	return d.DialAndSend(m)
}
