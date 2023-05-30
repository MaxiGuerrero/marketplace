package shared

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	DbConnection string
} 

func GetEnv() *env{
	if os.Getenv("GO_ENV") == "develop" || os.Getenv("GO_ENV") == "" {
		loadDotEnv()
	}
	return &env{
		DbConnection: os.Getenv("DB_CONNECTION"),
	}
}


func loadDotEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}