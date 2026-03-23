package server

import (
	"fmt"
	"log"
	"ms-images/pkg/dotenv"
)

func Init() {
	log.Printf("Init Server configuration")
	log.Printf("HOST: %s | PORT: %s", dotenv.GetEnv("MS_IMAGES_SERVER_URL"), dotenv.GetEnv("IMAGES_PORT"))
	r := NewRouter()
	r.Run(fmt.Sprintf("%s:%s", dotenv.GetEnv("MS_IMAGES_SERVER_URL"), dotenv.GetEnv("IMAGES_PORT")))
	log.Print("-----------------------------")
}
