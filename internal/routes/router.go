package routes

import (
	"FAMPAY/internal/handlers"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/v1")
	api.Use(CORSMiddleware())
	api.Get("/GetData", handlers.GetYoutubeDataPaginated)
}

func CORSMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Method() == fiber.MethodOptions {
			return
		}

		c.Next()
	}
}
