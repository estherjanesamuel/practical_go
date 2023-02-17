package interactor

import (
	models "RecipeApi/domain/model"
	"RecipeApi/usecase/presenter"
	"RecipeApi/usecase/repository"
)

type recipeInteractor struct {
	RecipeRepository repository.RecipeRepository
	RecipePresenter presenter.RecipePresenter
	DBRepository repository.DBRepository
}

type RecipeInteractor interface {
	Get(r []*models.Recipe) ([]*models.Recipe, error)
	// Create(r *models.Recipe) (*models.Recipe, error)
}

func NewRecipeInteractor(r repository.RecipeRepository, p presenter.RecipePresenter, d repository.DBRepository) RecipeInteractor {
	return &recipeInteractor{r, p, d}
}

func (rs *recipeInteractor) Get(r []*models.Recipe) ([]*models.Recipe, error) {
	r, err := rs.RecipeRepository.FindAll(r)
	if err != nil {
		return nil, err	
	}

	return rs.RecipePresenter.ResponseUser(r), nil

}