package responses

import "github.com/alexandrelamberty/macellum-api-server/pkg/domain"

func ListProvidersSuccessResponse(data *[]domain.Provider) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Providers successfully listed",
		Data:    data,
	}
}

func CreateProviderSuccessResponse(data *domain.Provider) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Provider successfully retrieved",
	}
}

func GetProviderSuccessResponse(data *domain.Provider) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Provider successfully retrieved",
	}
}

func UpdateProviderSuccessResponse(data *domain.Provider) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Provider successfully updated",
		Data:    data,
	}
}

func DeleteProviderSuccessResponse() ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Provider successfully deleted",
		Data:    nil,
	}
}

func ProviderErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: "Error: " + msg,
		Data:    nil,
	}
}
