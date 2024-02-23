package responses

import "github.com/alexandrelamberty/macellum-api-server/pkg/domain"

func ListCustomersSuccessResponse(data *[]domain.Customer) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Customers successfully listed",
		Data:    data,
	}
}

func CreateCustomerSuccessResponse(data *domain.Customer) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Customer successfully retrieved",
	}
}

func GetCustomerSuccessResponse(data *domain.Customer) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Data:    data,
		Message: "Customer successfully retrieved",
	}
}

func UpdateCustomerSuccessResponse(data *domain.Customer) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Customer successfully updated",
		Data:    data,
	}
}

func DeleteCustomerSuccessResponse() ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Customer successfully deleted",
		Data:    nil,
	}
}

func CustomerErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: "Error: " + msg,
		Data:    nil,
	}
}
