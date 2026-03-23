package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"ms-interactive-brokers/pkg/dotenv"
	"os"
	"time"
)

var db *gorm.DB

func InitDB() {
	log.Print("Init DB config")

	var host = dotenv.GetEnv("POSTGRES_HOST")
	var user = dotenv.GetEnv("POSTGRES_USER")
	var password = dotenv.GetEnv("POSTGRES_PASSWORD")
	var dbName = dotenv.GetEnv("POSTGRES_DB")
	var port = dotenv.GetEnv("POSTGRES_PORT")

	log.Printf("Initiating DB connection - host=%s user=%s dbname=%s port=%s",
		host,
		user,
		dbName,
		port)

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbName,
		port,
	)

	success, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}, // disables implicit prepared statement usage
	})

	if err != nil {
		log.Panicf("Error during DB setup %v", err)
	}

	db = success

	// Ensure the `uuid-ossp` extension is installed
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	log.Printf("Connected to DB %s", dotenv.GetEnv("NEWGO_DB"))
	log.Print("-----------------------------")
}

func GetDB() *gorm.DB {
	return db
}
