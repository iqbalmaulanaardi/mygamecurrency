package models

import (
	"errors"
	"time"
)

type Currency struct {
	CurrencyID uint64    `json:"currency_id" gorm:"primary_key;serial;"`
	Name       string    `json:"name" gorm:"unique;"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
}

func (c *Currency) Validate() error {
	if c.Name == "" {
		return errors.New("Invalid Param(s)")
	}
	return nil
}
