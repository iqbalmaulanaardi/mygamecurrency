package repository

import "github.com/iqbalmaulanaardi/mygamecurrency/models"

func (s *Service) InsertCurrency(currency models.Currency) (result models.Currency, err error) {
	err = s.DB.Model(models.Currency{}).Create(&currency).Error
	return
}
func (s *Service) GetAllCurrency() (result []models.Currency, err error) {
	err = s.DB.Model(models.Currency{}).Find(&result).Error
	return
}
