package routes

import (
	"rest/app/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// books
	books := e.Group("/books")
	books.POST("", handlers.CreateBook)
	books.GET("/:id", handlers.FindBook)
	books.GET("", handlers.ListBooks)
	books.PUT("/:id", handlers.UpdateBook)
	books.DELETE("/:id", handlers.DeleteBook)
}
