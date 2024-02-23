package main

import (
	"log"

	"github.com/alexandrelamberty/macellum-api-server/internal/routes"
	"github.com/alexandrelamberty/macellum-api-server/pkg/infrastructure"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Application entrypoint
func main() {
	// Environments variables
	// With Docker environment variables are loaded into the container.
	// With Go we load the local .env file
	// _, exist := os.LookupEnv("DATABASE_URI")
	// if !exist {
	// 	err := godotenv.Load(".env")
	// 	if err != nil {
	// 		log.Fatalf("Error loading .env file")
	// 	}
	// }

	// Set up database
	infrastructure.SetupDatabase()

	// Fiber Application
	app := fiber.New()

	// Cors
	app.Use(cors.New())

	// Logging
	app.Use("/", logger.New(logger.Config{
		Format:     "${cyan}[${time}] [${ip}]:${port} ${white}${pid} ${red}${status} - ${blue}${method} ${white}${path}\n",
		TimeFormat: "2006-01-02T15:04:05-0700",
		TimeZone:   "Europe/Brussels",
	}))

	// Repositories

	calendarRepo := repository.NewCalendarRepository(infrastructure.DB)
	cartRepo := repository.NewCartRepository(infrastructure.DB)
	categoryRepo := repository.NewCategoryRepository(infrastructure.DB)
	customerRepo := repository.NewCustomerRepository(infrastructure.DB)
	orderRepo := repository.NewOrderRepository(infrastructure.DB)
	productRepo := repository.NewProductRepository(infrastructure.DB)
	providerRepo := repository.NewProviderRepository(infrastructure.DB)
	userRepo := repository.NewUserRepository(infrastructure.DB)

	// Services
	calendarService := service.NewCalendarService(calendarRepo)
	cartService := service.NewCartService(cartRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	customerService := service.NewCustomerService(customerRepo)
	orderService := service.NewOrderService(orderRepo)
	productService := service.NewProductService(productRepo)
	providerService := service.NewProviderService(providerRepo)
	userService := service.NewUserService(userRepo)

	// app.Post("/users", func(c *fiber.Ctx) error {
	// 	var user domain.User
	// 	if err := c.BodyParser(&user); err != nil {
	// 		return err
	// 	}

	// 	if err := userService.CreateUser(&user); err != nil {
	// 		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	// 	}

	// 	return c.JSON(user)
	// })

	// app.Get("/users/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")
	// 	user, err := userService.GetUserByID(id)
	// 	if err != nil {
	// 		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	// 	}

	// 	return c.JSON(user)
	// })

	// app.Get("/users", func(c *fiber.Ctx) error {
	// 	users, err := userService.ListUsers()
	// 	if err != nil {
	// 		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	// 	}
	// 	return c.JSON(users)
	// })

	// Routes
	api := app.Group("/")
	routes.CalendarRouter(api, calendarService)
	routes.CartRouter(api, cartService)
	routes.CategoryRouter(api, categoryService)
	routes.CustomerRouter(api, customerService)
	routes.OrderRouter(api, orderService)
	routes.ProductRouter(api, productService)
	routes.ProviderRouter(api, providerService)
	routes.UserRouter(api, userService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Macellum API v0.0.1")
	})

	// Start the server
	log.Fatal(app.Listen(":3333"))
}
