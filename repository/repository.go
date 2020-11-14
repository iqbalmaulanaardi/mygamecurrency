package repository

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mygamecurrency/models"
	"github.com/jinzhu/gorm"
	"go.elastic.co/apm/module/apmgorm"
	_ "go.elastic.co/apm/module/apmgorm/dialects/postgres"
)

type Service struct {
	DB *gorm.DB
}

func Configure() (service Service, err error) {
	DbConfig := "host=" + "34.101.102.138" + " port=" + "5432" + " user=" + "postgres" + " dbname=" + "mygamecurrency" + " password=" + "postgres" + " sslmode=disable"
	if service.DB, err = apmgorm.Open("postgres", DbConfig); err != nil {
		return
	}
	ctx := context.Background()
	service.DB = apmgorm.WithContext(ctx, service.DB)
	sqlDB := service.DB.DB()
	sqlDB.SetMaxIdleConns(5)

	return
}
func GinHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("service", s)
		c.Next()
	}
}
func (s *Service) AutoMigrate() {
	s.DB.AutoMigrate(models.Currency{}, models.Conversion{})
	s.DB.Model(&models.Conversion{}).AddForeignKey("currency_id_from", "currencies(currency_id)", "RESTRICT", "RESTRICT")
	s.DB.Model(&models.Conversion{}).AddForeignKey("currency_id_to", "currencies(currency_id)", "RESTRICT", "RESTRICT")

}
