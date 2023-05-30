package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port int
}

var app *fiber.App

func CreateServer(port int) Server{
    return Server{Port: 8080}
}

func (server *Server) StartServer(routes *[]Route){
    app = fiber.New()
    registerRoutes(routes)
    var error = app.Listen(fmt.Sprintf(":%v", server.Port))
    if error != nil {
        log.Fatalln("Error to start server: ", error)
    }
}

func (server *Server) StopServer(){
    var error = app.Shutdown();
    if error != nil {
        log.Fatalln("Error to stop server: ", error)
    }
}