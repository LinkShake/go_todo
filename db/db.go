package database

import (
	"os"

	"github.com/LinkShake/go_todo/schema"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = Connect()

func Connect() *gorm.DB {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}
	connStr := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}
	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&schema.User{}, &schema.Todo{})
	if err != nil {
		panic(err)
	}
}