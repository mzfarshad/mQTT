package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/mzfarshad/MQTT-test/broker"
	"github.com/mzfarshad/MQTT-test/models"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %s", err)
	}
}

func main() {
	if err := models.ConnectPostgres(); err != nil {
		// panic("failed to connect database")
		log.Println("failed to connect database")
	} else {
		log.Println("successfully connected to database...")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	brokerClient, err := broker.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer brokerClient.Disconnect(2 * time.Second)

	err = brokerClient.Subscribe(broker.TopicRegisterCar, new(saveCar).WithBroker(brokerClient).Consume)
	if err != nil {
		log.Fatal(err)
	}

	<-c
}

// SaveCar is a consumer to save cars added to broker.TopicRegisterCar topic.
type saveCar struct {
	broker broker.Client
}

func (s *saveCar) WithBroker(c broker.Client) *saveCar {
	if s == nil {
		return new(saveCar).WithBroker(c)
	}
	s.broker = c
	return s
}

func (s *saveCar) Consume(msg []byte) {
	log.Printf(">>> consuming new message on %q...", broker.TopicRegisterCar.String())
	car := new(models.Car)
	err := json.Unmarshal(msg, &car)
	if err != nil {
		log.Printf("failed decoding JSON: %v", err)
		return
	}
	if err := car.Create(context.Background()); err != nil {
		log.Printf("failed to save car in database: %v", err)
		return
	}
	carMsg, _ := json.Marshal(car)
	err = s.broker.Publish(broker.TopicGetCars, carMsg)
	if err != nil {
		return
	}
}
