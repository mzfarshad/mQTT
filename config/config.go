package config

import (
	"log"
	"sync"
)

var config model
var once sync.Once

type Config interface {
	Postgres() *postgres
	Mqtt() *mqtt
}

func Get() Config {
	once.Do(
		func() {
			// Postgres
			psql, err := new(postgres).fromEnv()
			if err != nil {
				log.Println(err)
			}
			config.postgres = *psql
			// Mqtt
			config.mqtt = *new(mqtt).fromEnv()

		},
	)
	return config
}

type model struct {
	postgres
	mqtt
}

func (m model) Postgres() *postgres {
	return &m.postgres
}

func (m model) Mqtt() *mqtt {
	return &m.mqtt
}
