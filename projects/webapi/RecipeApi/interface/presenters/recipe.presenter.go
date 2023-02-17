package presenters

import (
	models "RecipeApi/domain/model"
	"RecipeApi/usecase/presenter"
)

type recipePresenter struct{}

func NewRecipePresenter() presenter.RecipePresenter {
	return &recipePresenter{}
}

func (rp *recipePresenter) ResponseUser(recipes []*models.Recipe) []*models.Recipe  {
	return recipes
}