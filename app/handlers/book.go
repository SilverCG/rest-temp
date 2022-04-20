package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/silvercg/rest-temp/app/models"
	"github.com/silvercg/rest-temp/app/store"
)

// Split the contract between the endpoints and the model
type bookResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    string    `json:"author"`
}

type bookRequest struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    string    `json:"author"`
}

func CreateBook(c echo.Context) error {

	book, err := MakeBookModel(c)
	if err != nil {
		return err
	}

	if err := store.CreateBook(c, book); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, MakeBookResponse(c, book))
}

func FindBook(c echo.Context) error {
	id := c.Param("id")
	bookID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	book := &models.Book{
		ModelV1: models.ModelV1{
			ID: bookID,
		},
	}
	if err := store.FindBook(c, book); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, MakeBookResponse(c, book))
}

func ListBooks(c echo.Context) error {
	var books []models.Book
	if err := store.ListBooks(c, &books); err != nil {
		return err
	}
	var bookResp []bookResponse
	for _, book := range books {
		bookResp = append(bookResp, MakeBookResponse(c, &book))
	}
	return c.JSON(http.StatusOK, bookResp)
}

func UpdateBook(c echo.Context) error {

	book, err := MakeBookModel(c)
	if err != nil {
		return err
	}

	if err := store.UpdateBook(c, book); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, MakeBookResponse(c, book))

}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	bookID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	book := &models.Book{
		ModelV1: models.ModelV1{
			ID: bookID,
		},
	}
	if err := store.DeleteBook(c, book); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, nil)

}

func MakeBookResponse(c echo.Context, book *models.Book) bookResponse {
	return bookResponse{
		ID:        book.ID,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
		Author:    book.Author,
	}
}

func MakeBookModel(c echo.Context) (*models.Book, error) {
	var bookReq bookRequest
	if err := c.Bind(&bookReq); err != nil {
		return nil, err
	}

	return &models.Book{
		ModelV1: models.ModelV1{
			ID: bookReq.ID,
		},
		Author: bookReq.Author,
	}, nil
}
