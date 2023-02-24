package main

import (
	"errors"
	"fiber1/business"
	"fiber1/config"
	"fiber1/model"
	"fiber1/presentation"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	db := config.ConnectMySql("root:password@(127.0.2.2:3306)/test1")
	if db == nil {
		panic(errors.New("cannot connect to database"))
	}

	guest := business.Guest{Db: db}
	_ = guest
	user := model.NewUser(db)
	_ = user.Migrate()

	app := fiber.New()

	app.Post(business.Guest_LoginPath, func(c *fiber.Ctx) error {
		in := business.Guest_LoginIn{}
		if err := c.BodyParser(&in); err != nil {
			return err
		}
		out := guest.Login(&in)
		return presentation.RenderRestApi(c, out, out.CommonOut)
	})
	presentation.HandleGet(app, business.Guest_LoginPath)
	
	app.Get(business.Guest_LoginPath, func(c *fiber.Ctx) error {return presentation.RenderHtml(c, business.Guest_LoginPath,nil)})

	app.Post(business.Guest_RegisterPath, func(c *fiber.Ctx) error {
		in := business.Guest_RegisterIn{}
		if err := c.BodyParser(&in); err != nil {
			return err
		}
		out := guest.Register(&in)
		return presentation.RenderRestApi(c, out, out.CommonOut)
	})
	presentation.HandleGet(app, business.Guest_RegisterPath)

	log.Fatal(app.Listen(":3000"))

}