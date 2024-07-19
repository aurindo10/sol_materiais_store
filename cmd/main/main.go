package main

import (
	"log"
	"os"

	"github.com/aurindo10/sol_store/internal/db"
	producthandler "github.com/aurindo10/sol_store/internal/handlers/productHandler"
	productrepository "github.com/aurindo10/sol_store/internal/repositories/productRepository"
	"github.com/aurindo10/sol_store/pkg/lib"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clerkSecretKey := os.Getenv("CLERK_SECRET_KEY")
	clerk.SetKey(clerkSecretKey)
	dbConnection := db.Connection()
	productRepo := productrepository.NewProductRepository(dbConnection)
	r.GET("/get_all_products", lib.ProtectedMiddlware(), producthandler.GetAllProducts(productRepo))
	r.POST("/create_products", lib.ProtectedMiddlware(), producthandler.CreateProduct(productRepo))
	r.GET("/test", producthandler.CreateProduct(productRepo))
	r.Run()
}
