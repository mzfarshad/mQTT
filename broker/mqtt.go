package broker

import (
	"log"
	"time"

	goMqtt "github.com/eclipse/paho.mqtt.golang"
)

type mqtt struct {
	client goMqtt.Client
}

func (m *mqtt) Subscribe(topic Topic, fn ConsumerFunc) error {
	token := m.client.Subscribe(
		topic.String(), 0,
		func(client goMqtt.Client, msg goMqtt.Message) {
			fn(msg.Payload())
		},
	)
	if token.Error() == nil {
		log.Printf("successfully subscribed to the topic of %q", topic.String())
	}
	return token.Error()
}

func (m *mqtt) Publish(topic Topic, msg []byte) error {
	token := m.client.Publish(topic.String(), 0, false, msg)
	if token.Wait() && token.Error() != nil {
		log.Printf("error publishing mqtt message : %v", token.Error())
		return token.Error()
	}
	log.Printf("successfully published to the topic of %q", topic.String())
	return nil
}

func (m *mqtt) Disconnect(after time.Duration) {
	m.client.Disconnect(uint(after.Microseconds()))
	log.Println("mqtt broker disconnected")
}
