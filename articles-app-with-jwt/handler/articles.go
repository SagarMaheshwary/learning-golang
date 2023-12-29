package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/database"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/model"
)

func GetArticles(c *fiber.Ctx) error {
	articles := new([]model.Article)

	database.DB.Preload("User").Find(&articles)

	return c.JSON(fiber.Map{
		"message": nil,
		"data": fiber.Map{
			"articles": articles,
		},
	})
}

func StoreArticle(c *fiber.Ctx) error {
	type ArticleInput struct {
		Title string
		Body  string
	}

	articleInput := new(ArticleInput)

	if err := c.BodyParser(&articleInput); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"message": "Unable to parse input."})
	}

	article := new(model.Article)

	article.Slug = strings.ReplaceAll(strings.ToLower(articleInput.Title), " ", "-")
	article.Title = articleInput.Title
	article.Body = articleInput.Body

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	article.UserId = uint(claims["user_id"].(float64))

	result := database.DB.Create(&article)

	if result.Error != nil {
		fmt.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": "Server error."})
	}

	database.DB.Preload("User").Find(&article)

	return c.JSON(fiber.Map{
		"message": "Store Article API",
		"data": fiber.Map{
			"article": article,
		},
	})
}

func GetArticle(c *fiber.Ctx) error {

	id := c.Params("id")

	article := new(model.Article)

	database.DB.Preload("User").First(&article, id)

	return c.JSON(fiber.Map{
		"message": nil,
		"data": fiber.Map{
			"article": article,
		},
	})
}

func UpdateArticle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update Article API"})
}

func DeleteArticle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Delete Article API"})
}
