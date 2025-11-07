package handlers

import (
    "github.com/gofiber/fiber/v2"
    "fibre-api/models"
    "fibre-api/services"
)

type UserHandler struct {
    Service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
    return &UserHandler{Service: service}
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
    users := h.Service.GetAllUsers()
    return c.JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
    }
    createdUser := h.Service.CreateUser(user)
    return c.Status(fiber.StatusCreated).JSON(createdUser)
}