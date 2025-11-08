package handlers

import (
    "github.com/gofiber/fiber/v2"
    "fibre-api/models"
    "fibre-api/services"
)

type BookHandler struct {
    Service services.BookService
}

func NewBookHandler(service services.BookService) *BookHandler {
    return &BookHandler{Service: service}
}

func (h *BookHandler) GetAllBooks(c *fiber.Ctx) error {
    books := h.Service.GetAllBooks()
    return c.JSON(books)
}

func (h BookHandler) CreateBooks(c *fiber.Ctx) error {
    var book models.Books
    if err := c.BodyParser(&book); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
    }
    createdbook := h.Service.CreateBook(book)
    return c.Status(fiber.StatusCreated).JSON(createdbook)
}