package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type postgres struct {
	Host     string
	Port     int
	User     string
	Pass     string
	Name     string
	TimeZone string
}

func (p *postgres) fromEnv() (*postgres, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return p, errors.New("invalid db port : " + err.Error())
	}
	p.Host = os.Getenv("DB_HOST")
	p.Port = port
	p.User = os.Getenv("DB_USER")
	p.Pass = os.Getenv("DB_PASS")
	p.Name = os.Getenv("DB_NAME")
	p.TimeZone = os.Getenv("DB_TIMEZONE")
	return p, nil
}
func (p *postgres) DNS() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		p.Host, p.User, p.Pass, p.Name, p.Port, p.TimeZone)
}
