package main

import (
	"fmt"
	"os"

	"github.com/silvercg/rest-temp/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// setup db and run migrations if any
	db := StartDatabase()

	// setup echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(DatabaseMiddleware(db))

	routes.InitRoutes(e)

	httpPort := os.Getenv("WEB_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func StartDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=% sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Failed to connect to database error: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("Failed to get database error: %v", err))
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(50)

	return db

}

func DatabaseMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
