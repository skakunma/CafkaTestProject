package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skakunma/CafkaTestProject/internal/config"
)

func main() {
	_, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	service := gin.Default()

	service.Run(":8080")
}
