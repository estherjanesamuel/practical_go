package controllers

import (
	"math"
	"recipe-api/pkg/common/models"
	"recipe-api/pkg/common/pagination"
	"recipe-api/pkg/common/validator"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// handler
type handler struct {
	DB *gorm.DB
}


// add recipe
func (h handler) AddRecipe(c *fiber.Ctx) error {
	body := validator.CreateRecipeRequestBody{}

	// parse body, attach to AddRecipeRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	errors := validator.ValidateStruct(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
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

// get all recipe with paginate

func paginate(value interface{}, pagination *pagination.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {  
    var totalRows int64 
    db.Model(value).Count(&totalRows)   

    pagination.TotalRows = totalRows    
    totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))    
    pagination.TotalPages = totalPages  

    return func(db *gorm.DB) *gorm.DB { 
        return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())   
    }   
}


// Get all recipe
func (h handler) GetRecipes(c *fiber.Ctx) error {
	
	var recipe []models.Recipe
	pagination := pagination.Pagination{}

	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")
	var sort = c.Query("sort", "")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	

	pagination.Page = intPage
	pagination.Limit = intLimit
	pagination.Sort = sort

	if result := h.DB.Scopes(paginate(recipe, &pagination, h.DB)).Find(&recipe).Error; result != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error())
	}

	pagination.Rows = recipe

	return c.Status(fiber.StatusOK).JSON(&pagination)
}


// update recipe
func (h handler) UpdateRecipe(c *fiber.Ctx) error {
	id := c.Params("id")
	body := validator.UpdateRecipeRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	errors := validator.ValidateStruct(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
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