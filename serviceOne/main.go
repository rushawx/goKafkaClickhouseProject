package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

type message struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	CreatedAt int64  `json:"created_at"`
}

func main() {

	producer, err := sarama.NewSyncProducer([]string{"kafka:29092"}, nil)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v\n", err)
	}
	defer producer.Close()

	router := gin.Default()

	router.GET("/click", func(c *gin.Context) {

		requestID := uuid.New().String()
		name := faker.FirstName()
		dtm := faker.Date()

		message := message{
			UserID:    requestID,
			Username:  name,
			Message:   fmt.Sprintf("click by %s at %v", name, dtm),
			CreatedAt: time.Now().Unix(),
		}

		bytes, err := json.Marshal(message)
		if err != nil {
			log.Printf("Error marshalling message: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal message"})
		}

		msg := &sarama.ProducerMessage{
			Topic: "clicks",
			Key:   sarama.StringEncoder(requestID),
			Value: sarama.ByteEncoder(bytes),
		}

		_, _, err = producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to produce message: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to produce message"})
		}

		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}
