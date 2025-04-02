package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skakunma/CafkaTestProject/internal/config"
	"github.com/skakunma/CafkaTestProject/internal/handlers"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	service := gin.Default()

	handlers.LoadHandlers(cfg, service)

	service.Run(":8080")
}
