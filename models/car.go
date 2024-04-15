package models

import (
	"context"
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

func (c *Car) Create(ctx context.Context) error {
	if c == nil {
		return errors.New("trying to create car model")
	}
	err := db.Debug().WithContext(ctx).Save(&c).Error

	if err != nil {
		return err
	}
	return nil
}

func (c Car) String() string {
	return fmt.Sprintf("Response save car :\n\tid: %d\n\tname: %s\n\tcompany: %s\n\tcolor: %s", c.ID, c.Name, c.Company, c.Color)
}
