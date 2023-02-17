package repository

import (
	models "RecipeApi/domain/model"
	"RecipeApi/usecase/repository"

	"gorm.io/gorm"
)

type recipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) repository.RecipeRepository  {
	return &recipeRepository{db}
}

func (rr *recipeRepository) FindAll(r []*models.Recipe) ([]*models.Recipe, error) {
	err := rr.db.Find(&r).Error

	if err != nil {
		return nil, err
	}

	return r, nil
}