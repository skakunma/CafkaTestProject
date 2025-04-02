package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/skakunma/CafkaTestProject/internal/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	UserForm struct {
		Username   string `json:"username" validate:"required"`
		Password   string `json:"password" validate:"required"`
		TelegramID string `json:"telegram_id" validate:"required"`
		Email      string `json:"email" validate:"required,email"`
	}
)

func Register(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		if !strings.HasPrefix(c.Request.Header.Get("Content-Type"), "application/json") {
			c.JSON(http.StatusBadRequest, "Conten type must be application/json")
			return
		}

		var infoUser UserForm
		err := c.ShouldBindJSON(&infoUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Body is nul")
			return
		}

		validate := validator.New()
		err = validate.Struct(infoUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, "JSON is not good")
			return
		}

		infoUser.Password = HashPassword(cfg, infoUser.Password)

		ctx := c.Request.Context()
		err = cfg.Store.CreateUser(ctx, infoUser.Username, infoUser.Password, infoUser.TelegramID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Problem with store")
			return
		}

		c.JSON(http.StatusCreated, "Account is created!")

	}
}

func HashPassword(cfg *config.Config, password string) string {
	hash := sha256.Sum256([]byte(password + cfg.Salt))
	return hex.EncodeToString(hash[:])
}
