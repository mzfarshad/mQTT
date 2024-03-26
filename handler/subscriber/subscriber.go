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
	saveCar    string = "cars/add-car"
	getCarByID string = "cars/get-car"
	allCars    string = "cars/all-cars"
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
		message := "Successfully saved car"
		token := client.Publish("response/save-car", 0, false, message)
		if token.Wait() && token.Error() != nil {
			log.Printf("error publishing MQTT message : %s\n", token.Error())
			return
		}
		fmt.Println("Saved car in database")

	case strings.HasPrefix(topic, getCarByID):
		log.Println(topic)
		car, err := models.FindCarByID(topic)
		if err != nil {
			log.Printf("error retrieveing car from database: %s\n", err)
			return
		}
		jsonCar, err := json.Marshal(car)
		if err != nil {
			log.Printf("failed encoding to jsonCar: %s\n", err)
			return
		}
		log.Println(string(jsonCar))
		token := client.Publish("response/car", 0, false, jsonCar)
		if token.Wait() && token.Error() != nil {
			log.Printf("error publishing MQTT message : %s\n", token.Error())
			return
		}

	case topic == allCars:
		log.Println(topic)
		cars, err := models.GetCars()
		if err != nil {
			log.Printf("error retrieveing cars from db: %s \n", err)
			return
		}
		jsonCars, err := json.Marshal(cars)
		if err != nil {
			log.Printf("failed encoding to jsonCars: %s \n", err)
			return
		}
		token := client.Publish("response/all-cars", 0, false, jsonCars)
		if token.Wait() && token.Error() != nil {
			log.Printf("error publishing MQTT message : %s\n", token.Error())
			return
		}
	}
}
