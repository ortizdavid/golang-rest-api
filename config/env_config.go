package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func ListenAddr() string {
	LoadDotEnv()
	return os.Getenv("APP_HOST")+":"+os.Getenv("APP_PORT")
}



