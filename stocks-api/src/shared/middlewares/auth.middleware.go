package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	config "marketplace/stocks-api/src/shared"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Middleware that is responsable to manage the authorization JWT token from a request.
func NewAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !config.GetConfig().Secure {
			return c.Next()
		}
		client := &http.Client{}
		authorization := c.Get("Authorization")
		urlSec := fmt.Sprintf("%v/token/validate",config.GetConfig().SecurityApi)
		request, err := http.NewRequestWithContext(c.Context(),"POST",urlSec,nil)
		if err != nil {
			log.Printf("Error on connect to Security API: %v", err.Error())
			return c.Status(500).JSON(config.InternalError("Error on connect to Security API"))
		}
		request.Header.Add("Authorization",authorization)
		res, err := client.Do(request)
		if err != nil {
			return c.Status(500).JSON(config.InternalError("Error on connect to Security API"))
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			body := &config.Response{}
			err := json.NewDecoder(res.Body).Decode(body)
			if err != nil {
				return c.Status(500).JSON(config.InternalError("Error on connect to Security API"))
			}
			return c.Status(res.StatusCode).JSON(body)
		}
		return c.Next() 
	}
}
