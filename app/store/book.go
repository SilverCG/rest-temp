package store

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/silvercg/rest-temp/app/models"
	"gorm.io/gorm"
)

// TODO replace echo.Context with an AppContext interface later.
func GetDBFromContext(c echo.Context) (*gorm.DB, error) {
	db := c.Get("db")
	if db.(*gorm.DB) != nil {
		return db.(*gorm.DB), nil
	}

	return nil, errors.New("No database was found in context")
}

func CreateBook(c echo.Context, book *models.Book) error {
	db, err := GetDBFromContext(c)
	if err != nil {
		return err
	}
	if err := db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func ListBooks(c echo.Context, books *[]models.Book) error {
	db, err := GetDBFromContext(c)
	if err != nil {
		return err
	}
	db.Find(books)
	return nil
}

func FindBook(c echo.Context, book *models.Book) error {
	db, err := GetDBFromContext(c)
	if err != nil {
		return err
	}
	db.Find(book)

	return nil
}

func UpdateBook(c echo.Context, book *models.Book) error {
	db, err := GetDBFromContext(c)
	if err != nil {
		return err
	}

	if err := db.Debug().Save(book).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(c echo.Context, book *models.Book) error {
	db, err := GetDBFromContext(c)
	if err != nil {
		return err
	}
	if err := db.Delete(book).Error; err != nil {
		return err
	}

	return nil
}
