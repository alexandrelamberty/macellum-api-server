package responses

import "github.com/alexandrelamberty/macellum-api-server/pkg/domain"

func ListCategoriesSuccessResponse(data *[]domain.Category) *[]domain.Category {
	return data
}

func CreateCategorySuccessResponse(data *domain.Category) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Category successfully retrieved",
	}
}

func GetCategorySuccessResponse(data *domain.Category) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Category successfully retrieved",
	}
}

func UpdateCategorySuccessResponse(data *domain.Category) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Category successfully updated",
		Data:    data,
	}
}

func DeleteCategorySuccessResponse() ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Category successfully deleted",
		Data:    nil,
	}
}

func CategoryErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: "Error: " + msg,
		Data:    nil,
	}
}
