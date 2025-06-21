package models

import "time"

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"user_id"`
	User      User        `gorm:"foreignKey:UserID" json:"user"`
	Name      string      `json:"name"`
	Phone     string      `json:"phone"`
	Address   string      `json:"address"`
	Note      string      `json:"note"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}
