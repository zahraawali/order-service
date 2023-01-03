package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zahraawali/order-servicec/pkg/db"
	"github.com/zahraawali/order-servicec/pkg/models"
)

func ListCart(c *gin.Context) {
	var carts []models.Cart

	db.DB.Find(&carts)

	c.JSON(http.StatusOK, gin.H{
		"carts": carts,
	})
}
