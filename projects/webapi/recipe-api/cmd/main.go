package main

import (
	"log"
	"recipe-api/pkg/common/config"
	"recipe-api/pkg/common/controllers"
	"recipe-api/pkg/common/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	controllers.RegisterRoutes(app, h)
	app.Listen(c.Port)
}