package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofor-little/env"
	_ "github.com/lib/pq"
)

// Not in use, just for reference
type User struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type Article struct {
	Id        string  `json:"id"`
	Slug      string  `json:"slug"`
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	UserId    string  `json:"user_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func main() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	db, err := getDatabaseConnection()

	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	app.Get("/articles", func(c *fiber.Ctx) error {
		return getArticles(c, db)
	})
	app.Post("/articles", func(c *fiber.Ctx) error {
		return storeArticle(c, db)
	})
	app.Get("/articles/:id", func(c *fiber.Ctx) error {
		return getArticle(c, db)
	})
	app.Put("/articles/:id", func(c *fiber.Ctx) error {
		return updateArticle(c, db)
	})
	app.Delete("/articles/:id", func(c *fiber.Ctx) error {
		return deleteArticle(c, db)
	})

	port := env.Get("APP_PORT", "8000")

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

}

func getArticles(c *fiber.Ctx, db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM articles")

	if err != nil {
		log.Fatalln(err)
		return c.JSON("An error occurred.")
	}

	//Docs: This releases any resources held by the rows no matter how the function returns.
	//Looping all the way through the rows also closes it implicitly, but it is better
	//to use defer to make sure rows is closed no matter what.
	defer rows.Close()

	var articles []Article

	for rows.Next() {
		var article Article

		rows.Scan(
			&article.Id,
			&article.Slug,
			&article.Title,
			&article.Body,
			&article.UserId,
			&article.CreatedAt,
			&article.UpdatedAt,
		)

		articles = append(articles, article)
	}

	return c.JSON(fiber.Map{
		"articles": articles,
	})
}

func storeArticle(c *fiber.Ctx, db *sql.DB) error {
	var article Article
	c.BodyParser(&article)

	article.Slug = strings.ReplaceAll(strings.ToLower(article.Title), " ", "-")

	err := db.QueryRow(
		"INSERT INTO articles (slug, title, body, user_id) VALUES ($1, $2, $3, $4) returning id, created_at",
		article.Slug,
		article.Title,
		article.Body,
		article.UserId,
	).Scan(&article.Id, &article.CreatedAt)

	if err != nil {
		log.Fatalln(err)
		return c.JSON("An error occurred.")
	}

	return c.JSON(fiber.Map{
		"message": "Created a new article.",
		"article": article,
	})
}

func getArticle(c *fiber.Ctx, db *sql.DB) error {
	id := c.Params("id")
	var article Article

	err := db.QueryRow("SELECT * FROM articles WHERE id = $1", id).Scan(
		&article.Id,
		&article.Slug,
		&article.Title,
		&article.Body,
		&article.UserId,
		&article.CreatedAt,
		&article.UpdatedAt,
	)

	if err != nil {
		log.Fatalln(err)
		return c.JSON("An error occurred.")
	}

	fmt.Println(article.Id, err)

	return c.JSON(fiber.Map{
		"article": article,
	})
}

func updateArticle(c *fiber.Ctx, db *sql.DB) error {
	id := c.Params("id")
	var article Article
	c.BodyParser(&article)

	article.Slug = strings.ReplaceAll(strings.ToLower(article.Title), " ", "-")

	err := db.QueryRow(
		"UPDATE articles SET slug = $1, title = $2, body = $3, updated_at = now() WHERE id = $4 returning id, user_id, created_at, updated_at",
		article.Slug,
		article.Title,
		article.Body,
		id,
	).Scan(&article.Id, &article.UserId, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		log.Fatalln(err)
		return c.JSON("An error occurred.")
	}

	return c.JSON(fiber.Map{
		"message": "Selected article has been updated.",
		"article": article,
	})
}

func deleteArticle(c *fiber.Ctx, db *sql.DB) error {
	id := c.Params("id")

	_, err := db.Exec("DELETE FROM articles WHERE id = $1", id)

	if err != nil {
		log.Fatalln(err)
		return c.JSON("An error occurred.")
	}

	return c.JSON(fiber.Map{
		"message": "Selected article has been deleted.",
	})
}

func getDatabaseConnection() (*sql.DB, error) {
	dbHost := env.Get("DB_HOST", "localhost")
	dbPort := env.Get("DB_PORT", "5432")
	dbDatabase := env.Get("DB_DATABASE", "articles_app")
	dbUsername := env.Get("DB_USERNAME", "postgres")
	dbPassword := env.Get("DB_PASSWORD", "password")
	dbSSLMode := env.Get("DB_SSLMODE", "disable")

	connectionURL := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbDatabase,
		dbSSLMode,
	)

	db, err := sql.Open("postgres", connectionURL)

	return db, err
}
