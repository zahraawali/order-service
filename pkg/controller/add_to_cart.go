package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zahraawali/order-servicec/pkg/db"
	"github.com/zahraawali/order-servicec/pkg/models"
	"github.com/zahraawali/order-servicec/protos/product"
)

func AddCart(c *gin.Context) {
	user_id, _ := c.Get("user_id")

	var cart models.Cart
	err := c.ShouldBind(&cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	_, err = db.Product.GetProduct(
		context.Background(),
		&product.PID{
			Id: int32(cart.ProductId),
		},
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "product_id is not valid.",
		})
		return
	}

	cart.UserId = fmt.Sprintf("%v", user_id)

	db.DB.Save(&cart)
}
