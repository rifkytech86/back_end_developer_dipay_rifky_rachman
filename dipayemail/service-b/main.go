package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"net/http"
	"os"
)

func sendEmailToKafka(subject string, body string) error {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	fmt.Println("starting producer")
	fmt.Println("starting producer", broker, topic)
	fmt.Println("starting producer")
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return err
	}
	defer producer.Close()

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(subject),
		Value:          []byte(body),
	}, nil)
	return err
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var inputData map[string]string
		// Decode JSON from request body
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&inputData); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			log.Println("Error decoding JSON:", err)
			return
		}

		email := inputData["email"]
		if email != "" {
			err := sendEmailToKafka("Welcome Dipay", "welcome, thanks joined with us")
			if err != nil {
				http.Error(w, "Error sending email", http.StatusInternalServerError)
				log.Println("Error sending email:", err)
				return
			}

			fmt.Fprintf(w, "Email request sent to Kafka")
		} else {
			http.Error(w, "Invalid request. Please provide 'email' and 'body' parameters.", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Invalid request method. Use POST.", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handleRequest)

	port := os.Getenv("SERVICE_B_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("HTTP server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
