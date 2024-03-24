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
	broker   string = "tcp://broker.emqx.io:1883"
	userName string = "emqx"
	password string = "public"
	clienID  string = "mqtt-test"
	saveCar  string = "cars/add-car"
	getCar   string = "cars/get-car"
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
	opts.SetUsername(userName)
	opts.SetPassword(password)
	opts.SetClientID(clienID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Subscribe(saveCar, 0, func(client mqtt.Client, msg mqtt.Message) {
		subscriber.CarSubscribe(client, msg)
	})
	client.Subscribe(getCar+"/#", 0, func(client mqtt.Client, msg mqtt.Message) {
		subscriber.CarSubscribe(client, msg)
	})

	<-c

	client.Disconnect(500)
	fmt.Println("Disconnect MQTT broker")
}
