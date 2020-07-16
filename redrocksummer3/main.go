package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redrocksummer3/rabbitmq"

	"redrocksummer3/model"
	"redrocksummer3/service"
)

func main()  {
	fmt.Println("Start now!")
	model.DbConn()
	go func() {
		rabbitmq.OpenConsumer()
	}()
	app := gin.Default()
	app.POST("/buy",service.Buy)
	app.Run(":8080")
}
