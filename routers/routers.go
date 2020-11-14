package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mygamecurrency/services"
)

func Configure(app *gin.Engine) {
	app.GET("/", services.Welcome)
	app.POST("/currency", services.AddCurrency)
	app.GET("/currencies", services.ListCurrency)
	app.POST("/conversion", services.AddConversion)
	app.POST("/calculate", services.Calculate)
}
