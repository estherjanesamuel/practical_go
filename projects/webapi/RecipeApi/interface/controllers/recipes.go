package controllers

import (
	"net/http"

	models "RecipeApi/domain/model"
	"RecipeApi/usecase/interactor"
)

type recipeController struct {
	recipeInteractor interactor.RecipeInteractor
}

type RecipeController interface {
	GetRecipes(c Context) error
}

func NewRecipeController(rs interactor.RecipeInteractor) RecipeController  {
	return &recipeController{rs}
}

func (rc *recipeController) GetRecipes(c Context) error {
	var r []*models.Recipe

	r, err := rc.recipeInteractor.Get(r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, r)
}

/*
import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// handler
type handler struct {
	DB *gorm.DB
}
type RecipeRequestBody struct {
	Name         string   `json:"name,omitempty"`
	Ingredients  []string `json:"ingredients,omitempty"`
	IsHalal      bool     `json:"isHalal"`
	IsVegetarian bool     `json:"isVegetarian"`
	Description  string   `json:"description,omitempty"`
	Rating       float64  `json:"rating"`
}

// add recipe
func (h handler) AddRecipe(c *fiber.Ctx) error {
	body := RecipeRequestBody{}

	// parse body, attach to AddRecipeRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	r := models.Recipe{
		Name:         body.Name,
		Ingredients:  body.Ingredients,
		IsHalal:      body.IsHalal,
		IsVegetarian: body.IsVegetarian,
		Description:  body.Description,
		Rating:       body.Rating,
	}
	
	// insert new db entry
	if result := h.DB.Create(&r); result.Error != nil {
		return fiber.NewError(fiber.StatusFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&r)
}

// delete recipe
func (h handler) DeleteRecipe(c *fiber.Ctx) error {
	id := c.Params("id")

	var recipe models.Recipe

	if result := h.DB.First(&recipe,id); result.Error != nil  {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	
	// delegate product from db
	h.DB.Delete(&recipe)
	return c.SendStatus(fiber.StatusOK)
}

// Get recipe by id

func (h handler) GetRecipe(c *fiber.Ctx) error {
	id := c.Params("id")
	var recipe models.Recipe
	
	if result := h.DB.First(&recipe,id); result.Error != nil  {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	
	return c.Status(fiber.StatusOK).JSON(&recipe)
}

// Get all recipe
func (h handler) GetRecipes(c *fiber.Ctx) error {
	var recipe []models.Recipe

	if result := h.DB.Find(&recipe); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&recipe)
}


// update recipe
func (h handler) UpdateRecipe(c *fiber.Ctx) error {
	id := c.Params("id")
	body := RecipeRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var recipe models.Recipe

	if result := h.DB.First(&recipe, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// populate body to update
	recipe.Name = body.Name
	recipe.Ingredients = body.Ingredients
	recipe.IsHalal = body.IsHalal
	recipe.IsVegetarian = body.IsVegetarian
	recipe.Description = body.Description
	recipe.Rating = body.Rating

	// save recipe
	h.DB.Save(&recipe)

	return c.Status(fiber.StatusOK).JSON(&recipe)
}


func RegisterRoutes(app *fiber.App, db *gorm.DB)  {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/recipes")
	routes.Post("/", h.AddRecipe)
	routes.Get("/", h.GetRecipes)
	routes.Get("/:id", h.GetRecipe)
	routes.Put("/:id", h.UpdateRecipe)
	routes.Delete("/:id", h.DeleteRecipe)

}

*/