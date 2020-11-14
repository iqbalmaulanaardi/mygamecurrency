package repository

import (
	"fmt"

	"github.com/iqbalmaulanaardi/mygamecurrency/models"
)

func (s *Service) FindCurrencyFromTo(c chan []models.Conversion, currencyIDFrom uint64, currencyIDTo uint64) {
	var (
		result      []models.Conversion
		err         error
		whereClause string
	)
	whereClause = fmt.Sprintf("(currency_id_to = %v AND currency_id_from = %v) OR (currency_id_to = %v AND currency_id_from = %v)", currencyIDFrom, currencyIDTo, currencyIDTo, currencyIDFrom)
	if err = s.DB.Model(models.Conversion{}).Where(whereClause).Find(&result).Error; err != nil {
		c <- []models.Conversion{}
	}
	c <- result
	return
}
func (s *Service) InsertConversion(conversion models.Conversion) (result models.Conversion, err error) {
	err = s.DB.Model(models.Conversion{}).Create(&conversion).Error
	return
}
