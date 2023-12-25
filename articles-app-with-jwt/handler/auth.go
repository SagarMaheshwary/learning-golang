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
	type RegisterInput struct {
		Name     string
		Email    string
		Password string
	}

	registerInput := new(RegisterInput)
	user := new(model.User)

	if err := c.BodyParser(&registerInput); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Unable to parse input."})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to hash password."})
	}

	user.Email = registerInput.Email
	user.Name = registerInput.Name
	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)

	if result.Error != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email already exists."})
	}

	token, err := createToken(user.Id, user.Email)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Server error."})
	}

	return c.JSON(fiber.Map{
		"message": "Registration successful.",
		"user":    user,
		"token":   token,
	})
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string
		Password string
	}

	loginInput := new(LoginInput)

	if err := c.BodyParser(&loginInput); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input."})
	}

	user := new(model.User)

	database.DB.Where("email = ?", loginInput.Email).Find(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials."})
	}

	token, err := createToken(user.Id, user.Email)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Server error."})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful.",
		"user":    user,
		"token":   token,
	})
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
