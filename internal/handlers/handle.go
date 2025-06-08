package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/zishang520/socket.io/v2/socket"
	"log"
	"ms-hooks/models"
	socket2 "ms-hooks/pkg/socket"
	"time"
)

func IniHandler(io *socket.Server) {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover2.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"name":    "MS Hooks",
			"version": "1.0.3",
		})
	})

	app.Get("/:id/*", func(c *fiber.Ctx) error {
		return process(c, io)
	})

	app.Post("/:id/*", func(c *fiber.Ctx) error {
		return process(c, io)
	})

	app.Delete("/:id/*", func(c *fiber.Ctx) error {
		return process(c, io)
	})

	app.Put("/:id/*", func(c *fiber.Ctx) error {
		return process(c, io)
	})

	err := app.Listen(":8081")
	if err != nil {
		log.Fatal(err)
	}
}

func process(c *fiber.Ctx, io *socket.Server) error {
	not, err := organizationData(c)
	if err != nil {
		return err
	}
	room := socket.Room(c.Params("id"))
	socket2.Emit(io, not, room)
	return c.JSON(map[string]string{"message": "receiver your data!", "status": "200"})
}

func organizationData(c *fiber.Ctx) (models.Notification, error) {
	var parse interface{}
	if len(c.Body()) != 0 {
		err := c.BodyParser(&parse)
		if err != nil {
			return models.Notification{}, err
		}
	}
	notification := models.Notification{
		Data:   parse,
		Header: c.GetReqHeaders(),
		Method: c.Method(),
		Time:   time.Now(),
		Path:   c.Path(),
	}
	return notification, nil
}
