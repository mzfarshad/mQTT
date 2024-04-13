package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"github.com/mzfarshad/MQTT-test/config"
	"github.com/mzfarshad/MQTT-test/handler/subscriber"
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
	}
	log.Println("successfully connected to database...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	opts := mqtt.NewClientOptions()
	mqttConfig := config.Get().Mqtt()
	opts.AddBroker(mqttConfig.BrokerAddress)
	opts.SetClientID(mqttConfig.ClientId)

	client := mqtt.NewClient(opts)
	disconnect := func(client mqtt.Client) {
		client.Disconnect(500)
		fmt.Println("mqtt broker disconnected")
	}
	defer disconnect(client)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Subscribe(subscriber.MqttTopicSaveCar.String(), 0, func(client mqtt.Client, msg mqtt.Message) {
		subscriber.CarSubscribe(client, msg)
	})
	client.Subscribe(subscriber.MqttToicGetCarByID.String()+"/#", 0, func(client mqtt.Client, msg mqtt.Message) {
		subscriber.CarSubscribe(client, msg)
	})
	client.Subscribe(subscriber.MqttTopicAllCars.String(), 0, func(client mqtt.Client, msg mqtt.Message) {
		subscriber.CarSubscribe(client, msg)
	})
	// log.Println("Hello World!")

	<-c

}
