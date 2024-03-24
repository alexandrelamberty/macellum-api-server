package handlers

import (
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllGroups(service service.GroupService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllGroups()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		return c.JSON(result)
	}
}

func CreateGroup(service service.GroupService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cart domain.Group
		// Parse the request body
		err := c.BodyParser(&cart)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Create the cart
		err = service.CreateGroup(&cart)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the created cart
		return c.JSON(&cart)
	}
}

func GetGroupByID(service service.GroupService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the cart from the service based on the parsed 'id'.Verify if the cart exists.
		group, err := service.GetGroupByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		return c.JSON(group)
	}
}

func UpdateGroup(service service.GroupService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var group domain.Group
		// Parse the request body
		err := c.BodyParser(&group)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Retrieve the cart from the service based on the parsed 'id'. Verify if the cart exists.
		_, err = service.GetGroupByID(group.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Update the cart
		err = service.UpdateGroup(&group)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the updated cart
		return c.JSON(&group)
	}
}

func DeleteGroup(service service.GroupService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the cart
		err = service.DeleteGroup(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		return c.JSON(true)
	}
}
