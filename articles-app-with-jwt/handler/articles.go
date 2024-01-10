package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/database"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/helpers"
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
	type RequestBody struct {
		Title string
		Body  string
	}

	requestBody := new(RequestBody)

	if err := c.BodyParser(&requestBody); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"message": "Invalid request.",
				"data":    nil,
			})
	}

	article := new(model.Article)

	article.Slug = helpers.CreateSlug(requestBody.Title)
	article.Title = requestBody.Title
	article.Body = requestBody.Body

	article.UserId = helpers.AuthUserId(c)

	if result := database.DB.Create(&article); result.Error != nil {
		fmt.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"message": "Server error.",
				"data":    nil,
			})
	}

	database.DB.Preload("User").Find(&article)

	return c.Status(fiber.StatusCreated).
		JSON(fiber.Map{
			"message": "Created a new article.",
			"data": fiber.Map{
				"article": article,
			},
		})
}

func GetArticle(c *fiber.Ctx) error {
	id := c.Params("id")

	article := new(model.Article)

	if result := database.DB.Preload("User").First(&article, id); result.Error != nil {
		fmt.Println(result.Error)
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{
				"message": "Article not found.",
				"data":    nil,
			})
	}

	return c.JSON(fiber.Map{
		"message": nil,
		"data": fiber.Map{
			"article": article,
		},
	})
}

func UpdateArticle(c *fiber.Ctx) error {
	type RequestBody struct {
		Title string
		Body  string
	}

	requestBody := new(RequestBody)

	if err := c.BodyParser(&requestBody); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"message": "Invalid request.",
				"data":    nil,
			})
	}

	id := c.Params("id")

	article := new(model.Article)

	if result := database.DB.First(&article, id); result.Error != nil {
		fmt.Println(result.Error)
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{
				"message": "Article not found.",
				"data":    nil,
			})
	}

	if article.UserId != helpers.AuthUserId(c) {
		return c.Status(fiber.StatusForbidden).
			JSON(fiber.Map{
				"message": "Unauthorized request.",
				"data":    nil,
			})
	}

	article.Slug = helpers.CreateSlug(requestBody.Title)
	article.Title = requestBody.Title
	article.Body = requestBody.Body

	if result := database.DB.Save(&article); result.Error != nil {
		fmt.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"message": "Server error.",
				"data":    nil,
			})
	}

	return c.JSON(fiber.Map{
		"message": "Article has been updated.",
		"data":    nil,
	})
}

func DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")

	article := new(model.Article)

	if result := database.DB.First(&article, id); result.Error != nil {
		fmt.Println(result.Error)
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{
				"message": "Article not found.",
				"data":    nil,
			})
	}

	if article.UserId != helpers.AuthUserId(c) {
		return c.Status(fiber.StatusForbidden).
			JSON(fiber.Map{
				"message": "Unauthorized request.",
				"data":    nil,
			})
	}

	if result := database.DB.Delete(article); result.Error != nil {
		fmt.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"message": "Server error.",
				"data":    nil,
			})
	}

	return c.JSON(fiber.Map{
		"message": "Article has been deleted.",
		"data":    nil,
	})
}
