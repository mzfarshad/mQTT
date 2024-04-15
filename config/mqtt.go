package config

import "os"

type mqtt struct {
	BrokerAddress string
	ClientId      string
}

func (m *mqtt) fromEnv() *mqtt {
	m.BrokerAddress = os.Getenv("MQTT_BROKER_ADDRESS")
	m.ClientId = os.Getenv("MQTT_CLIENT_ID")
	return m
}
