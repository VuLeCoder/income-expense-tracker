package main

import (
	"github.com/gin-gonic/gin"
	"server/configs"
	"server/databases"
	"server/models"
	"server/routers"
)

func main() {
	databases.ConnectDb()
	err := databases.DB.AutoMigrate(models.User{})
	if err != nil {
		return
	}

	router := gin.Default()

	configs.Cors(router)
	routers.Init(router)

	router.Run(":8080")
}