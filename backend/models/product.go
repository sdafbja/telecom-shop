package models

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	CategoryID  uint    `json:"category_id"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
	IsDeleted   bool    `gorm:"default:false" json:"is_deleted"`

}
