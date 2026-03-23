package main

import (
	"ms-images/db"
	"ms-images/models"
	"ms-images/pkg/dotenv"
	"ms-images/server"
)

func main() {
	dotenv.InitDotenv()
	dbSetup()
	server.Init()
}

func dbSetup() {
	db.Init()
	models.InitFileTypeModel()
	models.InitFileModel()
}
