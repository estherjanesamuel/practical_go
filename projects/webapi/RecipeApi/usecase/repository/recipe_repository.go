package repository

import models "RecipeApi/domain/model"

type RecipeRepository interface {
	FindAll(r []*models.Recipe) ([]*models.Recipe, error)
}