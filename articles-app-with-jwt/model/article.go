package model

import "time"

type Article struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	User      User
}
