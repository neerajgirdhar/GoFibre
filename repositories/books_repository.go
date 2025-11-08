package repositories

import "fibre-api/models"

// BookRepository defines the interface for book data interactions
type BooksRepository interface {
    GetAllBooks() []models.Books
    CreateBook(books models.Books) models.Books
}

type bookRepository struct {
    books []models.Books
}

func NewBookRepository() BooksRepository {
    return &bookRepository{
        books: []models.Books{},
    }
}

func (r *bookRepository) GetAllBooks() []models.Books {
    return r.books
}

func (r *bookRepository) CreateBook(newbook models.Books) models.Books {
    newbook.ID = len(r.books) + 1
    r.books = append(r.books, newbook)
    return newbook
}