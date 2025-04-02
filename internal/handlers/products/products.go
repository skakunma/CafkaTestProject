package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/skakunma/CafkaTestProject/internal/config"
	"github.com/skakunma/CafkaTestProject/internal/storage"
	"net/http"
	"strings"
)

type Product struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Categories  []int   `json:"categories"`
	Description string  `json:"description"`
}

func CreateProduct(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.Header.Get("Content-Type"), "application/json") {
			c.JSON(http.StatusBadRequest, "Content type must be application/json")
			return
		}

		var infoProduct Product

		err := c.ShouldBindJSON(&infoProduct)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Problem with parsing json")
			return
		}

		v := validator.New()
		err = v.Struct(infoProduct)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Problem in JSON")
			return
		}
		ctx := c.Request.Context()

		err = cfg.Store.CreateProduct(ctx, infoProduct.Name, infoProduct.Price, infoProduct.Categories, infoProduct.Description)
		if err != nil {
			if !errors.Is(storage.ErrNotFound, err) {
				c.JSON(http.StatusInternalServerError, "Service problem")
				return
			}
			c.JSON(http.StatusBadRequest, "Catigories is not found")
			return
		}

		c.JSON(http.StatusCreated, "Successful!")
		return
	}

}
