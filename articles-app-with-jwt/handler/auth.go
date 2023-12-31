package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/config"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/database"
	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/model"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	type RequestBody struct {
		Name     string
		Email    string
		Password string
	}

	requestBody := new(RequestBody)
	user := new(model.User)

	if err := c.BodyParser(&requestBody); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"message": "Invalid request.",
				"data":    nil,
			})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"message": "Server error.",
				"data":    nil,
			})
	}

	user.Email = requestBody.Email
	user.Name = requestBody.Name
	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"message": "Email already exists.",
				"data":    nil,
			})
	}

	token, err := createToken(user.Id, user.Email)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"message": "Server error.",
				"data":    nil,
			})
	}

	return c.JSON(fiber.Map{
		"message": "Registration successful.",
		"data": fiber.Map{
			"user":  user,
			"token": token,
		},
	})
}

func Login(c *fiber.Ctx) error {
	type RequestBody struct {
		Email    string
		Password string
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

	user := new(model.User)

	database.DB.Where("email = ?", requestBody.Email).Find(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{
				"message": "Invalid credentials.",
				"data":    nil,
			})
	}

	token, err := createToken(user.Id, user.Email)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": "Server error."})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful.",
		"data": fiber.Map{
			"user":  user,
			"token": token,
		},
	})
}

func UserProfile(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	user := new(model.User)
	database.DB.First(&user, claims["user_id"])

	return c.JSON(fiber.Map{
		"message": "User Profile API",
		"data": fiber.Map{
			"user": user,
		},
	})
}

func createToken(id uint, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	jwtKey := config.GetJWTConfig().Secret

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // @TODO: use expiry from JWT Config

	return token.SignedString([]byte(jwtKey))
}
