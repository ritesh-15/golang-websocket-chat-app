package database

import (
	"log"

	"github.com/ritesh-15/websocket-advanced/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{
		Conn: db,
	}
}

var DB *Database

func InitDatabase() {
	db, err := gorm.Open(mysql.Open(config.DATABASE_URL), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connecting to database", err)
		return
	}

	DB = NewDatabase(db)
	log.Println("Database connection established...")
}
