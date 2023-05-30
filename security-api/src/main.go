package main

import (
	"fmt"
	authentication "marketplace/security-api/src/authentication/infrastructure"
	server "marketplace/security-api/src/server"
	shared "marketplace/security-api/src/shared"
)

var routes []server.Route

func main(){
	a := shared.GetEnv()
	fmt.Println(a.DbConnection)
	instance := server.CreateServer(8080)
	routes = append(routes, authentication.GetRoutes()...)
	instance.StartServer(&routes)
}