package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skakunma/CafkaTestProject/internal/config"
	"github.com/skakunma/CafkaTestProject/internal/handlers/users"
)

func LoadHandlers(cfg *config.Config, c *gin.Engine) {
	c.POST("/register/", handlers.Register(cfg))
	c.POST("/login/")
}
