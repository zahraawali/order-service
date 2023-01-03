package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zahraawali/order-servicec/pkg/controller"
	"github.com/zahraawali/order-servicec/pkg/db"
	"github.com/zahraawali/order-servicec/pkg/middleware"
)

func init() {
	db.Connect()
	db.GrpcConnection()
}

func main() {
	router := gin.Default()

	router.GET("/cart", middleware.AuthMiddleware(), controller.ListCart)
	router.POST("/cart", middleware.AuthMiddleware(), controller.AddCart)
	router.DELETE("/cart", middleware.AuthMiddleware(), controller.EmptyCart)
	router.DELETE("/cart/{product_id}", middleware.AuthMiddleware(), controller.DeleteCart)

	if err := router.Run(":9000"); err != nil {
		log.Fatalln("Cannot listen on port 9000")
	}
}
