package database

import (
	"os"

	"github.com/LinkShake/go_todo/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connStr = os.Getenv("DATABASE_URL")

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
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