package router

import (
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/handler"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)
	api.Get("/profile", middleware.Authenticated(), handler.UserProfile)

	api.Get("/articles", handler.GetArticles)
	api.Post("/articles", middleware.Authenticated(), handler.StoreArticle)
	api.Get("/articles/:id", handler.GetArticle)
	api.Put("/articles/:id", middleware.Authenticated(), handler.UpdateArticle)
	api.Delete("/articles/:id", middleware.Authenticated(), handler.DeleteArticle)
}
