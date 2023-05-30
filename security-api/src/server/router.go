package server

import (
	"github.com/gofiber/fiber/v2"
)

type Route struct {
    Method string
    Path string
    Handler func(*fiber.Ctx) error
}

func registerRoutes(routes *[]Route){
    // Register routes
    for _, route := range *routes {
        app.Add(route.Method,route.Path,route.Handler)
    }
}
