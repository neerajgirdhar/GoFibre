package routes

import (
    "github.com/gofiber/fiber/v2"
    "fibre-api/handlers"
)

func UserRoutes(app *fiber.App, handler *handlers.UserHandler) {
    app.Get("/users", handler.GetAllUsers)
    app.Post("/users", handler.CreateUser)
}