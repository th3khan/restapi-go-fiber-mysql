package models

import "time"

type Product struct {
	Id           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	SerialNumber string    `json:"serial_number"`
	CreatedAt    time.Time `json:"created_at"`
}
