package validator

import "github.com/go-playground/validator"

var validate = validator.New()


type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type UpdateRecipeRequestBody struct {
	Name         string   `json:"name,omitempty"`
	Ingredients  []string `json:"ingredients,omitempty"`
	IsHalal      bool     `json:"isHalal"`
	IsVegetarian bool     `json:"isVegetarian"`
	Description  string   `json:"description,omitempty"`
	Rating       float64  `json:"rating"`
}


type CreateRecipeRequestBody struct {
	Name         string   `json:"name" validate:"required"`
	Ingredients  []string `json:"ingredients" validate:"required"`
	IsHalal      bool     `json:"isHalal,omitempty"`
	IsVegetarian bool     `json:"isVegetarian,omitempty"`
	Description  string   `json:"description" validate:"required"`
	Rating       float64  `json:"rating,omitempty"`
}