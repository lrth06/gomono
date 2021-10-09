package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lrth06/gomono/config"
	"github.com/lrth06/gomono/routes"
)

func main() {
	port := os.Getenv("PORT")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(recover.New())
	app.Use(logger.New())
	config.ConnDB("gomono")

	routes.SetupRoutes(app)

	app.Listen(":" + port)
}
