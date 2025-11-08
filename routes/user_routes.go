package routes

import (
    "github.com/gofiber/fiber/v2"
    "fibre-api/handlers"
)

func BookRoutes(app *fiber.App, handler *handlers.BookHandler) {
    app.Get("/books", handler.GetAllBooks)
    app.Post("/books", handler.CreateBooks)
}