package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/skakunma/CafkaTestProject/internal/config"
	"github.com/skakunma/CafkaTestProject/internal/jwtAuth"
	"net/http"
)

var handsNeedAuth = map[string]struct{}{
	"/balance/":   {},
	"/buy/:id":    {},
	"/products/":  {},
	"/user/:id":   {},
	"/me/":        {},
	"/me/edit/":   {},
	"/categorys/": {},
}

func AuthMiddleware(config *config.Config) gin.HandlerFunc {
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

		var claims jwtAuth.Claims
		jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtAuth.SecretKEY), nil
		})
		if err != nil || !jwtToken.Valid {
			c.JSON(http.StatusUnauthorized, "Token is broken")
			return
		}

		c.Set("user", claims)
		c.Next()
	}

}
