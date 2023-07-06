package shared

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Configurations that need the server.
type config struct {
	DbConnection string
	CostAlgorithmic int
	Database string
	JWTSecret []byte
	UrlApi string
	Port int
} 

// Get variable configurations with its respective type that has been setted via environment variables or from .env file.
// Those variables is use along of the system.
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

// Get variables from .env files.
func loadDotEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}

// Get URL API string either from env variable or a default URL.
func getUrlApi() string{
	urlApi, ok := os.LookupEnv("URL_API")
	if !ok {
		return fmt.Sprintf("http://localhost:%d",getPort())
	}
	return urlApi
}

// Get PORT API number either from env variable or a default PORT.
func getPort() int{
	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		return 8080
	}
	port , _ := strconv.Atoi(portStr)
	return port
}