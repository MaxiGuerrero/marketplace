package shared

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	DbConnection string
	CostAlgorithmic int
	Database string
	JWTSecret []byte
	UrlApi string
	Port int
} 

func GetConfig() *config{
	if os.Getenv("GO_ENV") == "develop" || os.Getenv("GO_ENV") == "" {
		loadDotEnv()
	}
	CostAlgorithmic , _ := strconv.Atoi(os.Getenv("COST_ALGORITHMIC"))
	
	return &config{
		DbConnection: os.Getenv("DB_CONNECTION"),
		CostAlgorithmic: CostAlgorithmic,
		Database: os.Getenv("DATABASE"),
		JWTSecret: []byte(os.Getenv("JWT_SECRET")),
		UrlApi: getUrlApi(),
		Port: getPort(),
	}
}


func loadDotEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}

func getUrlApi() string{
	urlApi, ok := os.LookupEnv("URL_API")
	if !ok {
		return fmt.Sprintf("http://localhost:%d",getPort())
	}
	return urlApi
}

func getPort() int{
	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		return 8080
	}
	port , _ := strconv.Atoi(portStr)
	return port
}