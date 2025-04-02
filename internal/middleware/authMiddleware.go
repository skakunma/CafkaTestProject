package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/skakunma/CafkaTestProject/internal/config"
	"net/http"
)

var handsNeedAuth = map[string]struct{}{
	"/balance/":  {},
	"/buy/:id":   {},
	"/products/": {},
	"/user/:id":  {},
	"/me/":       {},
	"/me/edit/":  {},
}

func AuthMiddleware(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exist := handsNeedAuth[c.Request.URL.Path]
		if !exist {
			c.Next()
			return
		}
		token, err := c.Cookie("jwt")
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, "Jwt is nul")
			return
		}

	}

}
