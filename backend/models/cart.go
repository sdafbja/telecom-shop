package models

type Cart struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"user_id"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`

	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}

