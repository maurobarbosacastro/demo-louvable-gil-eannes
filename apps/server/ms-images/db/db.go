package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"ms-images/pkg/dotenv"
	"os"
	"time"
)

var db *gorm.DB

func Init() {
	log.Print("Init DB config")
	log.Printf("Initiating DB connection - host=%s user=%s dbname=%s port=%s",
		dotenv.GetEnv("POSTGRES_HOST"),
		dotenv.GetEnv("POSTGRES_USER"),
		dotenv.GetEnv("POSTGRES_PASSWORD"),
		dotenv.GetEnv("POSTGRES_PORT"))

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dotenv.GetEnv("POSTGRES_HOST"),
		dotenv.GetEnv("POSTGRES_USER"),
		dotenv.GetEnv("POSTGRES_PASSWORD"),
		dotenv.GetEnv("IMAGES_DB"),
		dotenv.GetEnv("POSTGRES_PORT"),
	)

	success, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}))

	if err != nil {
		log.Panicf("Error during DB setup %v", err)
	}

	log.Print("Connecting to DB with the schema defined in configuration...")
	success, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})

	db = success
	log.Printf("Connected to DB %s", dotenv.GetEnv("IMAGES_DB"))
	log.Print("-----------------------------")
}

func GetDB() *gorm.DB {
	return db
}
