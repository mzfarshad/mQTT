package broker

import (
	"fmt"
	"time"

	goMqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mzfarshad/MQTT-test/config"
)

type ConsumerFunc func(message []byte)

type Client interface {
	Subscribe(topic Topic, fn ConsumerFunc) error
	Publish(topic Topic, msg []byte) error
	Disconnect(after time.Duration)
}

func NewClient() (Client, error) {
	opts := goMqtt.NewClientOptions()
	mqttConfig := config.Get().Mqtt()
	opts.AddBroker(mqttConfig.BrokerAddress)
	opts.SetClientID(mqttConfig.ClientId)

	client := goMqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed make mqtt client: %v", token.Error())
	}
	return &mqtt{
		client: client,
	}, nil
}
