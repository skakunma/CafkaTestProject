package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skakunma/CafkaTestProject/internal/config"
	handlers2 "github.com/skakunma/CafkaTestProject/internal/handlers/products"
	"github.com/skakunma/CafkaTestProject/internal/handlers/users"
	"github.com/skakunma/CafkaTestProject/internal/middleware"
)

func LoadHandlers(cfg *config.Config, c *gin.Engine) {
	c.Use(middleware.AuthMiddleware(cfg))
	c.POST("/register/", handlers.Register(cfg))
	c.POST("/login/", handlers.Login(cfg))
	c.POST("/product/", handlers2.CreateProduct(cfg))
}
