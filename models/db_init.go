package models

import (
	"github.com/mzfarshad/MQTT-test/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectPostgres() error {
	var err error
	psql := config.Get().Postgres()
	db, err = gorm.Open(postgres.Open(psql.DNS()), &gorm.Config{})
	if err != nil {
		return err
	}
	return migrate(db)
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Car{},
	)
}
