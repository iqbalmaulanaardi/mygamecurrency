package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mygamecurrency/models"
	"github.com/iqbalmaulanaardi/mygamecurrency/repository"
)

func ListCurrency(c *gin.Context) {
	var (
		err    error
		result []models.Currency
	)
	service := c.MustGet("service").(*repository.Service)
	if result, err = service.GetAllCurrency(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
	return
}
