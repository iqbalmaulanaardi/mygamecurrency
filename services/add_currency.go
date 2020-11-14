package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mygamecurrency/models"
	"github.com/iqbalmaulanaardi/mygamecurrency/repository"
)

func AddCurrency(c *gin.Context) {
	var (
		err      error
		currency models.Currency
	)
	if err = c.Bind(&currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Param(s)"})
		return
	}
	if err = currency.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	service := c.MustGet("service").(*repository.Service)
	if _, err = service.InsertCurrency(currency); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
	return
}
