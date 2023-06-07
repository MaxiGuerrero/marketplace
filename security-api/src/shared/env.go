package shared

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	DbConnection string
	CostAlgorithmic int
} 

func GetConfig() *config{
	if os.Getenv("GO_ENV") == "develop" || os.Getenv("GO_ENV") == "" {
		loadDotEnv()
	}
	CostAlgorithmic , _ := strconv.Atoi(os.Getenv("COST_ALGORITHMIC"))
	return &config{
		DbConnection: os.Getenv("DB_CONNECTION"),
		CostAlgorithmic: CostAlgorithmic,
	}
}


func loadDotEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}