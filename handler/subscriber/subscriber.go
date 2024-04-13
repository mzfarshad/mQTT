package subscriber

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mzfarshad/MQTT-test/models"
)

type MqttTopic string

func (m MqttTopic) String() string {
	return string(m)
}

const (
	MqttTopicSaveCar   MqttTopic = "cars/add-car"
	MqttToicGetCarByID MqttTopic = "cars/get-car"
	MqttTopicAllCars   MqttTopic = "cars/all-cars"
)

func CarSubscribe(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()

	switch {
	case topic == MqttTopicSaveCar.String():
		log.Println(topic)
		car := new(models.Car)
		err := json.Unmarshal(msg.Payload(), &car)
		if err != nil {
			log.Printf("failed decoding JSON: %v", err)
			return
		}
		if err := car.Create(); err != nil {
			log.Printf("failed to save car in database: %v", err)
			return
		}
		message := "Successfully saved car"
		token := client.Publish("response/save-car", 0, false, message)
		if token.Wait() && token.Error() != nil {
			log.Printf("error publishing MQTT message : %v", token.Error())
			return
		}
		fmt.Println("Saved car in database")

	case strings.HasPrefix(topic, MqttToicGetCarByID.String()):
		log.Println(topic)
		getID := strings.Split(topic, "/")
		id, err := strconv.Atoi(getID[len(getID)-1])
		if err != nil {
			log.Printf("failed get id from topic : %v", err)
			response := "failed get id from topic"
			token := client.Publish("response/car", 0, false, response)
			if token.Wait() && token.Error() != nil {
				log.Printf("error publishing MQTT message : %v", token.Error())
				return
			}
			return
		}
		car, err := models.FindCarByID(id)
		if err != nil {
			log.Printf("error retrieveing car from database: %v", err)
			response := fmt.Sprintf("not found car by id : %d please try again...", id)
			token := client.Publish("response/car", 0, false, response)
			if token.Wait() && token.Error() != nil {
				log.Printf("error publishing MQTT message : %v", token.Error())
				return
			}
			return
		}
		if len(car) < 1 {
			response := fmt.Sprintf("id:%d is not exist", id)
			log.Print(response)
			token := client.Publish("response/car", 0, false, response)
			if token.Wait() && token.Error() != nil {
				log.Printf("error publishing MQTT message : %v", token.Error())
				return
			}
			return
		}
		jsonCar, err := json.Marshal(car)
		if err != nil {
			log.Printf("failed encoding to jsonCar: %v", err)
			return
		}
		log.Println(string(jsonCar))
		token := client.Publish("response/car", 0, false, jsonCar)
		if token.Wait() && token.Error() != nil {
			log.Printf("error publishing MQTT message : %v", token.Error())
			return
		}

	case topic == MqttTopicAllCars.String():
		log.Println(topic)
		cars, err := models.GetCars()
		if err != nil {
			log.Printf("error retrieveing cars from db: %v", err)
			return
		}
		jsonCars, err := json.Marshal(cars)
		if err != nil {
			log.Printf("failed encoding to jsonCars: %v", err)
			return
		}
		token := client.Publish("response/all-cars", 0, false, jsonCars)
		if token.Wait() && token.Error() != nil {
			log.Printf("error publishing MQTT message : %v", token.Error())
			return
		}
	}
}
