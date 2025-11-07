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
    userRepo := repositories.NewUserRepository()
    userService := services.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    // Setup routes
    routes.UserRoutes(app, userHandler)

    // Start the server
    app.Listen(":3000")
}