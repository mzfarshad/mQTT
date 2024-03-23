package config

import (
	"log"
	"sync"
)

var config model
var once sync.Once

type Config interface {
	Postgres() *postgres
}

func Get() Config {
	once.Do(
		func() {
			psql, err := new(postgres).fromEnv()
			if err != nil {
				log.Println(err)
			}
			config.postgres = *psql
		},
	)
	return config
}

type model struct {
	postgres
}

func (m model) Postgres() *postgres {
	return &m.postgres
}
