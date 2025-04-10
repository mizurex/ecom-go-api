package main

import (
	"ecom/config"
	"ecom/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	fmt.Println("starting")
	handler.SetDB(db)
	router := gin.Default()
	router.POST("/user", handler.AddUserHandler)
	router.Run("localhost:9090")

}
