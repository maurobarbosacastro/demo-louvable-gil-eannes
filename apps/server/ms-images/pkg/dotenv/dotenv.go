package dotenv

import (
	"github.com/joho/godotenv"
	"os"
)

func InitDotenv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
