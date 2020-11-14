package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mygamecurrency/models"
	"github.com/iqbalmaulanaardi/mygamecurrency/repository"
	"github.com/iqbalmaulanaardi/mygamecurrency/requests"
)

func Calculate(c *gin.Context) {
	var (
		err              error
		calculateRequest requests.CalculateRequest
		checkConversion  []models.Conversion
		result           float64
	)
	if err = c.Bind(&calculateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Param(s)"})
		return
	}
	if err = calculateRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	service := c.MustGet("service").(*repository.Service)
	chanConversion := make(chan []models.Conversion)
	go service.FindCurrencyFromTo(chanConversion, calculateRequest.CurrencyIDFrom, calculateRequest.CurrencyIDTo)
	checkConversion = <-chanConversion
	if len(checkConversion) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid Currency ID"})
		return
	}
	if checkConversion[0].CurrencyIDFrom == calculateRequest.CurrencyIDFrom {
		result = checkConversion[0].Rate * calculateRequest.Amount
		c.JSON(http.StatusOK, gin.H{"result": result})
		return
	} else {
		result = calculateRequest.Amount / checkConversion[0].Rate
		c.JSON(http.StatusOK, gin.H{"result": result})
		return
	}

}
