package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zahraawali/order-servicec/pkg/models"
	"github.com/zahraawali/order-servicec/protos/auth"
	"github.com/zahraawali/order-servicec/protos/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var DB *gorm.DB
var Auth auth.AuthServiceClient
var Product product.ProductServiceClient

func Connect() {
	var err error
	DB, err = gorm.Open("mysql", "devops:Passw0rd!@tcp(127.0.0.1:3300)/cloud_compute?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.Cart{})
}

func GrpcConnection() {
	authConn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error Dial: ", err)
	}

	prodConn, err := grpc.Dial("127.0.0.1:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error Dial: ", err)
	}

	Auth = auth.NewAuthServiceClient(authConn)
	Product = product.NewProductServiceClient(prodConn)
}
