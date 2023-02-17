package presenter

import models "RecipeApi/domain/model"

type RecipePresenter interface {
	ResponseUser(r []*models.Recipe) []*models.Recipe
}