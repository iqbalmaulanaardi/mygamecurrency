package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mygamecurrency/models"
	"github.com/iqbalmaulanaardi/mygamecurrency/repository"
)

func AddConversion(c *gin.Context) {
	var (
		err             error
		conversion      models.Conversion
		checkConversion []models.Conversion
	)
	chanConversion := make(chan []models.Conversion)
	if err = c.Bind(&conversion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Param(s)"})
		return
	}
	if err = conversion.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	service := c.MustGet("service").(*repository.Service)
	//check if from to exist
	go service.FindCurrencyFromTo(chanConversion, conversion.CurrencyIDFrom, conversion.CurrencyIDTo)
	checkConversion = <-chanConversion
	fmt.Println(len(checkConversion))
	if len(checkConversion) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Conversion ID already exist"})
		return
	}
	if _, err = service.InsertConversion(conversion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
	return
}
