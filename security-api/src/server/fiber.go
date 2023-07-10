package server

import (
	"fmt"
	"log"
	docs "marketplace/security-api/src/docs"
	responses "marketplace/security-api/src/shared"
	"syscall"
	"unsafe"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// Responsable to implement the API server logical to start or stop the server.
// This server implement the library "Fiber".
type Server struct {
	Port int
    App *fiber.App
}

// Create a instance of the server setting the port and if the swagger doc must be expose.
func CreateServer(port int,activateDocs bool) *Server{
	var mask uintptr
	
	// Get the current CPU affinity of the process
	if _, _, err := syscall.RawSyscall(syscall.SYS_SCHED_GETAFFINITY, 0, uintptr(unsafe.Sizeof(mask)), uintptr(unsafe.Pointer(&mask))); err != 0 {
        panic("Failed to get CPU affinity:")
	}
	fmt.Println("Current CPU affinity:", mask)

	// Set the new CPU affinity
	mask = 24
	if _, _, err := syscall.RawSyscall(syscall.SYS_SCHED_SETAFFINITY, 0, uintptr(unsafe.Sizeof(mask)), uintptr(unsafe.Pointer(&mask))); err != 0 {
        panic("Failed to set CPU affinity:")
	}
	fmt.Println("New CPU affinity:", mask)
    app := fiber.New(fiber.Config{
        ErrorHandler: errorHandler,
    })
    app.Use(LogRequest)
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
    return &Server{Port: port, App: app}
}

// Run server on port configurated.
func (server *Server) StartServer(){
    var error = server.App.Listen(fmt.Sprintf(":%v", server.Port))
    if error != nil {
        log.Fatalln("Error to start server: ", error)
    }
}

// Stop server running.
func (server *Server) StopServer(){
    var error = server.App.Shutdown();
    if error != nil {
        log.Fatalln("Error to stop server: ", error)
    }
}

// Logical about handle internal error when a request throw a Panic.
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

// Little middleware that log a request when start and end it.
func LogRequest(c *fiber.Ctx) error {
    log.Printf("Starting request: %v",c.Context().ID())
    err := c.Next()
    log.Printf("Request completed: %v",c.Context().ID())
    return err
}