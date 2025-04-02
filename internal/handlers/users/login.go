package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/skakunma/CafkaTestProject/internal/config"
	"github.com/skakunma/CafkaTestProject/internal/jwtAuth"
	"github.com/skakunma/CafkaTestProject/internal/storage"
	"net/http"
	"strings"
)

func Login(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.Header.Get("Content-Type"), "application/json") {
			c.JSON(http.StatusBadRequest, "Content-Type must be application/json")
			return
		}

		var infoLogin UserForm

		err := c.ShouldBindJSON(&infoLogin)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Problem in body request")
			return
		}

		v := validator.New()
		err = v.Struct(infoLogin)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Problem with JSON")
			return
		}

		infoLogin.Password = HashPassword(cfg, infoLogin.Password)

		ctx := c.Request.Context()
		password, err := cfg.Store.GetPasswordFromUsername(ctx, infoLogin.Username)
		if err != nil {
			if !errors.Is(err, storage.ErrNotFound) {
				c.JSON(http.StatusInternalServerError, "Problem in service")
				return
			}
			c.JSON(http.StatusBadRequest, "Username is not found")
			return
		}

		if password != infoLogin.Password {
			c.JSON(http.StatusBadRequest, "Password is not correct")
			return
		}

		userId, err := cfg.Store.GetIdFromUsername(ctx, infoLogin.Username)
		if err != nil {
			if !errors.Is(err, storage.ErrNotFound) {
				c.JSON(http.StatusInternalServerError, "service problem")
				return
			}
			c.JSON(http.StatusBadRequest, "username is not found")
			return
		}

		token, err := jwtAuth.BuildJWTString(userId)

		c.SetCookie("jwt", token, 3600, "/", "", false, false)

		return
	}

}
