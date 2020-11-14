package requests

import "errors"

type CalculateRequest struct {
	CurrencyIDFrom uint64  `json:"currency_id_from"`
	CurrencyIDTo   uint64  `json:"currency_id_to"`
	Amount         float64 `json:"amount"`
}

func (c *CalculateRequest) Validate() error {
	if c.CurrencyIDTo == 0 || c.CurrencyIDFrom == 0 || c.Amount == 0.0 {
		return errors.New("Invalid Param(s)")
	}
	return nil
}
