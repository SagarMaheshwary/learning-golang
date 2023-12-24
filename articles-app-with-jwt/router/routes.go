package router

import (
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)
	api.Get("/profile", handler.UserProfile)

	api.Get("/articles", handler.GetArticles)
	api.Post("/articles", handler.StoreArticle)
	api.Get("/articles/:id", handler.GetArticle)
	api.Put("/articles/:id", handler.UpdateArticle)
	api.Delete("/articles/:id", handler.DeleteArticle)
}
