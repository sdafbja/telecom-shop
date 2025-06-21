package models

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

    ProductID uint   `json:"product_id"`
    UserID    uint   `json:"user_id"`
    Rating    int    `json:"rating"`
    Comment   string `json:"comment"`
    User      User   `json:"user" gorm:"foreignKey:UserID"`
}

