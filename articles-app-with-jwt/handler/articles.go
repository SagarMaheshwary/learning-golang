package handler

import "github.com/gofiber/fiber/v2"

func GetArticles(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Articles Listing API"})
}

func StoreArticle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Store Article API"})
}

func GetArticle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Article Details API"})
}

func UpdateArticle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update Article API"})
}

func DeleteArticle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Delete Article API"})
}
