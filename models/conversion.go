package models

import (
	"errors"
	"time"
)

type Conversion struct {
	CurrencyIDFrom uint64    `json:"currency_id_from"`
	CurrencyIDTo   uint64    `json:"currency_id_to"`
	Rate           float64   `json:"rate"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
}

func (c *Conversion) Validate() error {
	if c.CurrencyIDFrom == 0 || c.CurrencyIDTo == 0 || c.Rate == 0.0 {
		return errors.New("Invalid Param(s)")
	}
	if c.CurrencyIDFrom == c.CurrencyIDTo {
		return errors.New("From and To Can not be the same")
	}
	return nil
}
