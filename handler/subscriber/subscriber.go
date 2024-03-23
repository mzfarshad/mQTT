package subscriber

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mzfarshad/MQTT-test/models"
)

const (
	saveCar string = "/cars/add-car"
	getCar  string = "/cars/get-car"
)

func CarSubscribe(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()

	switch {
	case topic == saveCar:
		log.Println(topic)
		car := new(models.Car)
		err := json.Unmarshal(msg.Payload(), &car)
		if err != nil {
			log.Printf("failed decoding JSON: %s\n", err)
			return
		}
		if err := car.Create(); err != nil {
			log.Printf("failed to save car in database: %s\n", err)
			return
		}
		fmt.Println("Saved car in database")
	case strings.HasPrefix(topic, getCar):
		log.Println(topic)
		car, err := models.FindCarByID(topic)
		if err != nil {
			log.Printf("error retrieveing car from database: %s\n", err)
			return
		}
		jsonCar, err := json.Marshal(car)
		if err != nil {
			log.Printf("failed encoding to json: %s\n", err)
			return
		}
		log.Println(string(jsonCar))
		responseTopic := fmt.Sprintf("/response%s", topic)
		log.Println(responseTopic)
		token := client.Publish(responseTopic, 0, false, jsonCar)
		if token.Wait() && token.Error() != nil {
			log.Printf("error publishing MQTT message : %s\n", token.Error())
			return
		}
	}
}
