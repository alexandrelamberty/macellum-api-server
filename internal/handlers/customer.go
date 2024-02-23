package handlers

import (
	"fmt"
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/requests"
	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func GetAllCustomers(service service.CustomerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllCustomers()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}
		return c.JSON(responses.ListCustomersSuccessResponse(result))
	}
}

func CreateCustomer(service service.CustomerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Request
		var request requests.CreateCustomer

		// Parse the request body
		err := c.BodyParser(&request)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}

		// Validate the request
		err = validate.Struct(request)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}

		// Map the request to the domain model
		customer := domain.Customer{
			Code:      request.Code,
			FirstName: request.FirstName,
			LastName:  request.LastName,
			BirthDate: request.BirthDate,
			Address: domain.Address{
				Street:  request.Address.Street,
				Number:  request.Address.Number,
				ZipCode: request.Address.ZipCode,
				City:    request.Address.City,
				State:   request.Address.State,
				Country: request.Address.Country,
			},
			Phone:    request.Phone,
			Email:    request.Email,
			Notes:    request.Notes,
			Status:   request.Status,
			Label:    request.Label,
			Priority: request.Priority,
		}

		// Create the customer
		err = service.CreateCustomer(&customer)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}

		// Return the created customer
		return c.JSON(responses.GetCustomerSuccessResponse(&customer))
	}
}

func GetCustomerByID(service service.CustomerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the customer from the service based on the parsed 'id'.Verify if the customer exists.
		customer, err := service.GetCustomerByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}
		fmt.Println(customer)

		return c.JSON(responses.GetCustomerSuccessResponse(customer))
	}
}

func UpdateCustomer(service service.CustomerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var customer domain.Customer
		// Parse the request body
		err := c.BodyParser(&customer)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}

		// Retrieve the customer from the service based on the parsed 'id'. Verify if the customer exists.
		_, err = service.GetCustomerByID(customer.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}

		// Update the customer
		err = service.UpdateCustomer(&customer)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}

		// Return the updated customer
		return c.JSON(responses.UpdateCustomerSuccessResponse(&customer))
	}
}

func DeleteCustomer(service service.CustomerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the customer
		err = service.DeleteCustomer(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CustomerErrorResponse(err.Error()))
		}

		// Return the deleted customer ID
		// FIXME: Return the deleted customer id
		return c.JSON(responses.DeleteCustomerSuccessResponse())
	}
}
