package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/config"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/database"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/model"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse input."})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to hash password."})
	}

	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)

	if result.Error != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already exists."})
	}

	token, err := createToken(user.Id, user.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to create token."})
	}

	return c.JSON(fiber.Map{
		"message": "Register API",
		"user":    user,
		"token":   token,
	})
}

func Login(c *fiber.Ctx) error {
	var user model.User

	database.DB.Find(&user)

	return c.JSON(fiber.Map{"message": "Login API"})
}

func UserProfile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "User Profile API"})
}

func createToken(id uint, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // @TODO: use expiry from JWT

	return token.SignedString([]byte(config.GetJWTConfig().Secret))
}
