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
	Database string
	UrlApi string
	Port int
	SecurityApi string
	Secure bool
	MqServerUrl string
} 

// Get variable configurations with its respective type that has been setted via environment variables or from .env file.
// Those variables is use along of the system.
func GetConfig() *config{
	if os.Getenv("GO_ENV") == "develop" || os.Getenv("GO_ENV") == "" {
		loadDotEnv()
	}
	return &config{
		DbConnection: os.Getenv("DB_CONNECTION"),
		Database: os.Getenv("DATABASE"),
		UrlApi: getUrlApi(),
		Port: getPort(),
		SecurityApi: os.Getenv("SECURITY_API_URL"),
		Secure: os.Getenv("SECURE") ==  "true",
		MqServerUrl: getMqServerUrl(),
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
		return 8081
	}
	port , _ := strconv.Atoi(portStr)
	return port
}

func getMqServerUrl() string {
	urlMq, ok := os.LookupEnv("MQ_SERVER_URL")
	if !ok {
		return "tcp://127.0.0.1:9000"
	}
	return urlMq
}