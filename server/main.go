package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lrth06/gomono/config"
	// "github.com/joho/godotenv"

)

func SetupRoutes ( app *fiber.App){
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, from GO 👋!")
	})
	app.Get("/ping", func(c *fiber.Ctx)error{
		return c.Status(200).SendString("Server is available and healthy")
	})
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) 
	})
}



func main() {
	app := fiber.New(fiber.Config{
		// IdleTimeout: idleTimeout,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(recover.New())
	SetupRoutes(app)
	config.Database()

	log.Fatal(app.Listen(":5000"))	

}
