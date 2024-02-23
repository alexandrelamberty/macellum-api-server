package routes

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service service.UserService) {
	app.Get("/users", handlers.GetAllUsers(service))
	app.Post("/users", handlers.CreateUser(service))
	app.Get("/users/:id", handlers.GetUserByID(service))
	app.Patch("/users/:id", handlers.UpdateUser(service))
	app.Delete("/users/:id", handlers.DeleteUser(service))
}
