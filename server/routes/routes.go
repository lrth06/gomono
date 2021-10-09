package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/gomono/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", handlers.Hello)
	app.Get("/ping", handlers.Ping)
	app.Get("/users", handlers.GetUsers)
	app.Get("/posts", handlers.GetPosts)
	app.Post("/users", handlers.CreateUser)
	app.Get("/users/:id", handlers.GetUserById)
	app.Delete("/users/:id", handlers.DeleteUser)
	app.Put("/users/:id", handlers.UpdateUser)
}
