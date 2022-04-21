package models

import "time"

type Order struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	ProductId int       `json:"product_id"`
	Product   Product   `gorm:"foreignkey:ProductId"`
	UserId    int       `json:"user_id"`
	User      User      `gorm:"foreignkey:UserId"`
	CreatedAt time.Time `json:"created_at"`
}
