package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zahraawali/order-servicec/pkg/db"
	"github.com/zahraawali/order-servicec/pkg/models"
)

type ID struct {
	ProductID uint `json:"product_id"`
}

func DeleteCart(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	user_id = fmt.Sprintf("%v", user_id)

	var product_id_str string = c.Param("product_id")
	product_id, err := strconv.ParseUint(product_id_str, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "product_id is not valid.",
		})
		return
	}

	db.DB.Delete(&models.Cart{}, "user_id = ? AND product_id = ?", user_id, product_id)
	c.JSON(http.StatusOK, gin.H{
		"status": "Ok",
	})
}
