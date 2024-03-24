package router

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func IncomeRouter(app fiber.Router, service service.IncomeService) {
	app.Get("/incomes", handlers.GetAllIncomes(service))
	app.Post("/incomes", handlers.CreateIncome(service))
	app.Get("/incomes/:id", handlers.GetIncomeByID(service))
	app.Patch("/incomes/:id", handlers.UpdateIncome(service))
	app.Delete("/incomes/:id", handlers.DeleteIncome(service))
}
