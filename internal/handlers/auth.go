package handlers

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/requests"
	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the request body
		var request *requests.RegisterUser
		err := c.BodyParser(&request)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Map the request to the domain model
		user := &domain.User{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Password:  request.Password,
		}

		// Register the user
		err = service.RegisterUser(user)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Return the registered user
		return c.JSON(responses.GetUserSuccessResponse(user))
	}
}

func LoginUser(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the request body
		var request requests.LoginUser
		err := c.BodyParser(&request)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Validate the request body
		err = validate.Struct(request)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Login the user with credentials
		user, err := service.LoginUser(request.Email, request.Password)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Return the logged user
		return c.JSON(responses.GetUserSuccessResponse(user))
	}
}
