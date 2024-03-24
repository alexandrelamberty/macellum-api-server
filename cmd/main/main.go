package main

import (
	"log"

	"github.com/alexandrelamberty/macellum-api-server/internal/database"
	"github.com/alexandrelamberty/macellum-api-server/internal/router"
	"github.com/alexandrelamberty/macellum-api-server/pkg/infrastructure"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Application entrypoint
func main() {

	// Set up database
	infrastructure.SetupDatabase()

	// Seed database
	database.SeedDatabase(infrastructure.DB)

	// Fiber Application
	app := fiber.New()

	// Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

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
	groupRepo := repository.NewGroupRepository(infrastructure.DB)
	incomeRepo := repository.NewIncomeRepository(infrastructure.DB)
	locationRepo := repository.NewLocationRepository(infrastructure.DB)
	limitRepo := repository.NewLimitRepository(infrastructure.DB)
	paymentMethodRepo := repository.NewPaymentMethodRepository(infrastructure.DB)
	customerRepo := repository.NewCustomerRepository(infrastructure.DB)
	orderRepo := repository.NewOrderRepository(infrastructure.DB)
	productRepo := repository.NewProductRepository(infrastructure.DB)
	providerRepo := repository.NewProviderRepository(infrastructure.DB)
	userRepo := repository.NewUserRepository(infrastructure.DB)

	// Services
	authService := service.NewAuthService(userRepo)
	calendarService := service.NewCalendarService(calendarRepo)
	cartService := service.NewCartService(cartRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	groupService := service.NewGroupService(groupRepo)
	incomeService := service.NewIncomeService(incomeRepo)
	locationService := service.NewLocationService(locationRepo)
	limitService := service.NewLimitService(limitRepo)
	paymentMethodService := service.NewPaymentMethodService(paymentMethodRepo)
	customerService := service.NewCustomerService(customerRepo)
	orderService := service.NewOrderService(orderRepo)
	productService := service.NewProductService(productRepo)
	providerService := service.NewProviderService(providerRepo)
	userService := service.NewUserService(userRepo)

	// Routes
	api := app.Group("/")
	router.AuthRouter(api, authService)
	router.CalendarRouter(api, calendarService)
	router.CartRouter(api, cartService)
	router.CategoryRouter(api, categoryService)
	router.GroupRouter(api, groupService)
	router.IncomeRouter(api, incomeService)
	router.LocationRouter(api, locationService)
	router.LimitRouter(api, limitService)
	router.PaymentMethodRouter(api, paymentMethodService)
	router.CustomerRouter(api, customerService)
	router.OrderRouter(api, orderService)
	router.ProductRouter(api, productService)
	router.ProviderRouter(api, providerService)
	router.UserRouter(api, userService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Macellum API v0.0.1")
	})

	// Start the server
	log.Fatal(app.Listen(":3333"))
}
