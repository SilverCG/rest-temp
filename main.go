package main

import (
	"fmt"
	"os"

	"github.com/silvercg/rest-temp/app/models"
	"github.com/silvercg/rest-temp/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// load env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("Could not load env file error: %s", err))
	}

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

	// connect to postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
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

	// set connection limits
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("Failed to get database error: %v", err))
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(50)

	// automigrate for now. TODO: add up/down style migrations later
	db.Migrator().DropTable(&models.Book{})
	db.AutoMigrate(&models.Book{})

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
