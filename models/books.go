package models
type Books struct {
    ID    int    `json:"id"`
    BookIsbn    int    `json:"bookisbn"`
    BookName  string `json:"bookname"`
    Author string `json:"bookauthor"`
}