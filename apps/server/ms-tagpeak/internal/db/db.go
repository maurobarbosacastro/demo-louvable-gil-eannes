package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"os"
	"time"
)

var db *gorm.DB

func InitDB() {
	logster.StartFuncLog()

	var host = dotenv.GetEnv("POSTGRES_HOST")
	var user = dotenv.GetEnv("POSTGRES_USER")
	var password = dotenv.GetEnv("POSTGRES_PASSWORD")
	var dbName = dotenv.GetEnv("TAGPEAK_DB")
	var port = dotenv.GetEnv("POSTGRES_PORT")

	logster.Info(fmt.Sprintf("Initiating DB connection - host=%s user=%s dbname=%s port=%s",
		host,
		user,
		dbName,
		port))

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbName,
		port,
	)

	success, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true,
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
		logster.Panic(err, "Error during DB setup")
	}

	db = success

	// Ensure the `uuid-ossp` extension is installed
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	logster.Info(fmt.Sprintf("Connected to DB %s", dbName))
	logster.EndFuncLog()
}

func GetDB() *gorm.DB {
	return db
}
