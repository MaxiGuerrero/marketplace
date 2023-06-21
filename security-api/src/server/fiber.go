package server

import (
	"fmt"
	"log"
	responses "marketplace/security-api/src/shared"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	Port int
    App *fiber.App
}

func CreateServer(port int) *Server{
    numGoroutines := runtime.NumCPU()
	fmt.Println("Number of CPUs:", numGoroutines)
	runtime.GOMAXPROCS(numGoroutines)
    app := fiber.New(fiber.Config{
        ErrorHandler: errorHandler,
    })
    app.Use(recover.New())
    return &Server{Port: 8080, App: app}
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

func errorHandler(ctx *fiber.Ctx, err error) error {
    // Status code defaults to 500
    code := fiber.StatusInternalServerError

    e := ctx.Status(code).JSON(responses.InternalError(err.Error()))

    if e != nil {
        return e
    }

    // Return from handler
    return nil
}