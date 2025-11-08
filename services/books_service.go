package services

import (
    "fibre-api/models"
    "fibre-api/repositories"
)

type BookService interface {
    GetAllBooks() []models.Books
    CreateBook(books models.Books) models.Books
}

type bookService struct {
    repo repositories.BooksRepository
}

func NewBookService(repo repositories.BooksRepository) BookService {
    return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() []models.Books {
    return s.repo.GetAllBooks()
}

func (s *bookService) CreateBook(books models.Books) models.Books {
    return s.repo.CreateBook(books)
}