package models

import "time"

type Cart struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	ProductId int       `json:"product_id"`
	UserId    string    `json:"user_id omit"`
	Total     uint      `json:"total"`
}
