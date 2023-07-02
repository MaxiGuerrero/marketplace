package server

import (
	"fmt"
	"log"
	docs "marketplace/security-api/src/docs"
	responses "marketplace/security-api/src/shared"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type Server struct {
	Port int
    App *fiber.App
}

func CreateServer(port int,activateDocs bool) *Server{
    numGoroutines := runtime.NumCPU()
	fmt.Println("Number of CPUs:", numGoroutines)
	runtime.GOMAXPROCS(numGoroutines)
    app := fiber.New(fiber.Config{
        ErrorHandler: errorHandler,
    })
    if(activateDocs){
        doc := docs.LoadDoc()
        app.Get("/docs/*", swagger.New(swagger.Config{
            URL:         "/swagger/doc.json",
            DeepLinking: false,
	    }))
        app.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
            return c.SendString(doc)
        })
    }
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