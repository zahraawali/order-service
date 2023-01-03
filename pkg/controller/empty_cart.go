package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zahraawali/order-servicec/pkg/db"
	"github.com/zahraawali/order-servicec/pkg/models"
)

func EmptyCart(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	user_id = fmt.Sprintf("%v", user_id)

	db.DB.Delete(models.Cart{}, "user_id = ?", user_id)

	c.JSON(http.StatusOK, gin.H{
		"status": "Ok",
	})
}
