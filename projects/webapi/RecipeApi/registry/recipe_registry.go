package registry

import (
	"RecipeApi/interface/controllers"
	"RecipeApi/interface/presenters"
	"RecipeApi/interface/repository"
	"RecipeApi/usecase/interactor"
)


func (r *registry) NewRecipeController() controllers.RecipeController {
	recipeInteractor := interactor.NewRecipeInteractor(
		repository.NewRecipeRepository(r.db),
		presenters.NewRecipePresenter(),
		repository.NewDBRepository(r.db),
	)

	return controllers.NewRecipeController(recipeInteractor) 
}
