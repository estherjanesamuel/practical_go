package router

import (
	"RecipeApi/interface/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, h controllers.AppController, cc controllers.Context) {

	routes := app.Group("/recipes")
	routes.Get("/", func(c *fiber.Ctx) error {return h.Recipe.GetRecipes(cc)})
	// routes.Post("/", h.)
	// routes.Get("/", func (c fiber.Ctx) error {return h.Recipe.GetRecipes(c)})
	// routes.Get("/:id", h.GetRecipe)
	// routes.Put("/:id", h.UpdateRecipe)
	// routes.Delete("/:id", h.DeleteRecipe)

}