package router

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GroupRouter(app fiber.Router, service service.GroupService) {
	app.Get("/groups", handlers.GetAllGroups(service))
	app.Post("/groups", handlers.CreateGroup(service))
	app.Get("/groups/:id", handlers.GetGroupByID(service))
	app.Patch("/groups/:id", handlers.UpdateGroup(service))
	app.Delete("/groups/:id", handlers.DeleteGroup(service))
}
