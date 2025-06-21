package models

// CartInput dùng để thêm sản phẩm vào giỏ hàng
type CartInput struct {
	ProductID uint `json:"product_id" example:"1"`
	Quantity  int  `json:"quantity" example:"2"`
}

// UpdateCartQuantityInput dùng để cập nhật số lượng giỏ hàng
type UpdateCartQuantityInput struct {
	Quantity int `json:"quantity" example:"3"`
}
