package handlers

import (
	"fmt"
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/requests"
	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllUsers()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}
		return c.JSON(responses.ListUsersSuccessResponse(result))
	}
}

func InviteUser(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request *requests.InviteUser
		// Parse the request body
		err := c.BodyParser(&request)
		if err != nil {
			fmt.Println("error body")
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Validate the request body
		err = validate.Struct(request)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// User
		user := &domain.User{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
		}

		// Create the user
		err = service.InviteUser(user)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Return the created user
		return c.JSON(responses.GetUserSuccessResponse(user))
	}
}

func GetUserByID(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the user from the service based on the parsed 'id'.Verify if the user exists.
		user, err := service.GetUserByID(uintID)
		if err != nil {
			c.Status(http.StatusNotFound)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		return c.JSON(responses.GetUserSuccessResponse(user))
	}
}

func UpdateUser(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user domain.User
		// Parse the request body
		err := c.BodyParser(&user)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Retrieve the user from the service based on the parsed 'id'. Verify if the user exists.
		_, err = service.GetUserByID(user.ID)
		if err != nil {
			c.Status(http.StatusNotFound)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Update the user
		err = service.UpdateUser(&user)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Return the updated user
		return c.JSON(responses.UpdateUserSuccessResponse(&user))
	}
}

func DeleteUser(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the user
		err = service.DeleteUser(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}

		// Return the deleted user ID
		// FIXME: Return the deleted user id
		return c.JSON(responses.DeleteUserSuccessResponse())
	}
}
