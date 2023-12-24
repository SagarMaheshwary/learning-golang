package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofor-little/env"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/config"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/database"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/router"
)

func main() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	database.Connect()

	app := fiber.New()

	router.InitRoutes(app)

	port := config.GetAppConfig().AppPort

	app.Listen(fmt.Sprintf(":%s", port))
}
