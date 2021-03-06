package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	conn "github.com/iqbalmaulanaardi/mygamecurrency/repository"
	"github.com/iqbalmaulanaardi/mygamecurrency/routers"
)

func main() {
	var (
		service conn.Service
		err     error
	)
	app := gin.New()
	port := ":8072"
	app.MaxMultipartMemory = 10
	app.Use(cors.Default())
	if service, err = conn.Configure(); err != nil {
		panic(err)
	} else {
		service.AutoMigrate()
		fmt.Println("Connected to db")
	}
	app.Use(conn.GinHandler(&service))
	routers.Configure(app)
	app.Run(port)
}
