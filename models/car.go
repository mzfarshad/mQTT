package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Color   string `json:"color"`
}

func (c *Car) Create() error {
	if c == nil {
		return errors.New("trying to create car model")
	}
	err := db.Debug().Save(&c).Error

	if err != nil {
		return err
	}
	return nil
}

func FindCarByID(id int) ([]Car, error) {
	var cars []Car
	if err := db.Where("id=?", id).Find(&cars).Error; err != nil {
		return nil, fmt.Errorf("not found car: %s", err)
	}
	return cars, nil
}

func GetCars() ([]Car, error) {
	var cars []Car
	if err := db.Debug().Find(&cars).Error; err != nil {
		return nil, fmt.Errorf("error retrieved from db : %s", err)
	}
	return cars, nil
}
