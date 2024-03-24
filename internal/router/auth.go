package router

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service service.AuthService) {
	auth := app.Group("/auth")
	auth.Post("/register", handlers.RegisterUser(service))
	auth.Post("/login", handlers.LoginUser(service))
}
