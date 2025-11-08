package main

import (
    "github.com/gofiber/fiber/v2"
    "fibre-api/handlers"
    "fibre-api/repositories"
    "fibre-api/routes"
    "fibre-api/services"
)

func main() {
    app := fiber.New()

    // Initialize repository, service, and handler
    bookRepo := repositories.NewBookRepository()
    bookService := services.NewBookService(bookRepo)
    bookHandler := handlers.NewBookHandler(bookService)

    // Setup routes
    routes.BookRoutes(app, bookHandler)

    // Start the server
    app.Listen(":3000")
}