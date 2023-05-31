package main

import (
	authentication "marketplace/security-api/src/authentication/infrastructure"
	s "marketplace/security-api/src/server"
)

func main(){
	server := s.CreateServer(8080)
	authentication.RegisterRoutes(server.App)
	server.StartServer()
}