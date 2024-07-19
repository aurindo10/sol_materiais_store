package db

import (
	"log"
	"os"

	"github.com/aurindo10/sol_store/internal/db/entities"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	err := godotenv.Load()
	connStr := os.Getenv("GOOSE_DBSTRING")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(entities.Products{})
	return db
}
