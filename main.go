package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"github.com/mzfarshad/MQTT-test/handler/subscriber"
	"github.com/mzfarshad/MQTT-test/models"
)

const (
	broker     string = "tcp://localhost:1885"
	clienID    string = "mqtt-test"
	saveCar    string = "cars/add-car"
	getCarByID string = "cars/get-car"
	allCars    string = "cars/all-cars"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %s", err)
	}
}
func main() {
	if err := models.ConnectPostgres(); err != nil {
		panic("failed to connect database")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clienID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Subscribe(saveCar, 0, func(client mqtt.Client, msg mqtt.Message) {
		subscriber.CarSubscribe(client, msg)
	})
	client.Subscribe(getCarByID+"/#", 0, func(client mqtt.Client, msg mqtt.Message) {
		subscriber.CarSubscribe(client, msg)
	})
	client.Subscribe(allCars, 0, func(client mqtt.Client, msg mqtt.Message) {
		subscriber.CarSubscribe(client, msg)
	})

	<-c

	client.Disconnect(500)
	fmt.Println("Disconnect MQTT broker")
}
