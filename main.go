package main

import (
	"github.com/jihadable/stockwise-be/database"
	"github.com/jihadable/stockwise-be/middlewares"
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/jihadable/stockwise-be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	api := app.Group("/api", middlewares.ErrorHandler())
	db := database.DB()

	err = db.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		panic(err)
	}

	routes.RegisterUserRoutes(api, db)
	routes.RegisterProductRoutes(api, db)

	err = app.Listen(":3000")
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
