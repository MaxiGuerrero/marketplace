package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port int
    App *fiber.App
}

func CreateServer(port int) *Server{
    return &Server{Port: 8080, App: fiber.New()}
}

func (server *Server) StartServer(){
    var error = server.App.Listen(fmt.Sprintf(":%v", server.Port))
    if error != nil {
        log.Fatalln("Error to start server: ", error)
    }
}

func (server *Server) StopServer(){
    var error = server.App.Shutdown();
    if error != nil {
        log.Fatalln("Error to stop server: ", error)
    }
}