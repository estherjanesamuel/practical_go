package main

import (
	"RecipeApi/config"
	"RecipeApi/infrastrcture/db"
	"RecipeApi/infrastrcture/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}

	db := db.Init(&c)
	app := fiber.New()
	// h controllers.AppController, cc controllers.Context
	router.RegisterRoutes(app, h)
	// controllers.RegisterRoutes(app, h)
	app.Listen(c.Port)
}