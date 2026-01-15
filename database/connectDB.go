package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	_ = godotenv.Load()

	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Database connection established")
			return
		} else {
			log.Printf("❌ Database connection failed: %v\n", err)
		}

		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("❌ Could not connect to the database after several attempts:", err)
	}
}
