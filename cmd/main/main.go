package main

import (
	"log"
	"os"

	"github.com/aurindo10/sol_store/internal/db"
	producthandler "github.com/aurindo10/sol_store/internal/handlers/productHandler"
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
	r.GET("/all_products", lib.ProtectedMiddlware(), producthandler.CreateProduct(dbConnection))
	r.GET("/test", producthandler.CreateProduct(dbConnection))
	r.Run()
}
