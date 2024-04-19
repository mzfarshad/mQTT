package config

import "os"

type mqtt struct {
	BrokerAddress string
	ClientId      string
	UserName      string
	Password      string
}

func (m *mqtt) fromEnv() *mqtt {
	m.BrokerAddress = os.Getenv("MQTT_BROKER_ADDRESS")
	m.ClientId = os.Getenv("MQTT_CLIENT_ID")
	m.UserName = os.Getenv("MQTT_USERNAME")
	m.Password = os.Getenv("MQTT_PASSWORD")
	return m
}
