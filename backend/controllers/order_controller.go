package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/config"
	"github.com/sdafbja/telecom-shop/models"
)

// 🧾 Tạo đơn hàng từ giỏ hàng và thông tin xác nhận (name, phone, address, note)
// CreateOrder godoc
// @Summary Tạo đơn hàng
// @Description Tạo đơn hàng từ giỏ hàng và thông tin xác nhận (tên, số điện thoại, địa chỉ, ghi chú)
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body struct{ name string; phone string; address string; note string } true "Thông tin xác nhận đơn hàng"
// @Success 201 {object} models.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders [post]
// @Security BearerAuth


func CreateOrder(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input struct {
		Name    string  `json:"name"`
		Phone   string  `json:"phone"`
		Address string  `json:"address"`
		Note    string  `json:"note"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu xác nhận không hợp lệ"})
		return
	}

	// Lấy giỏ hàng
	var cartItems []models.Cart
	if err := config.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil || len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Giỏ hàng trống hoặc không thể truy cập"})
		return
	}

	var total float64
	var orderItems []models.OrderItem
	for _, item := range cartItems {
		subtotal := float64(item.Quantity) * item.Product.Price
		total += subtotal

		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})
	}

	// Tạo đơn hàng
	order := models.Order{
		UserID:    userID,
		Name:      input.Name,
		Phone:     input.Phone,
		Address:   input.Address,
		Note:      input.Note,
		Total:     total,
		Status:    "pending",
		CreatedAt: time.Now(),
		Items:     orderItems,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo đơn hàng"})
		return
	}

	config.DB.Where("user_id = ?", userID).Delete(&models.Cart{})

	c.JSON(http.StatusCreated, order)
}



// 👤 Xem đơn hàng của chính mình
// GetMyOrders godoc
// @Summary Xem đơn hàng cá nhân
// @Description Trả về danh sách đơn hàng của người dùng hiện tại
// @Tags Orders
// @Produce json
// @Success 200 {array} models.Order
// @Failure 500 {object} map[string]string
// @Router /orders/my [get]
// @Security BearerAuth
func GetMyOrders(c *gin.Context) {
	userID := c.GetUint("user_id")
	var orders []models.Order

	err := config.DB.
		Preload("Items.Product").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy đơn hàng"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// 🛠️ Admin: Xem tất cả đơn hàng
// GetAllOrders godoc
// @Summary Quản trị: Xem tất cả đơn hàng
// @Description Trả về danh sách toàn bộ đơn hàng (chỉ dành cho Admin)
// @Tags Orders
// @Produce json
// @Success 200 {array} models.Order
// @Failure 500 {object} map[string]string
// @Router /admin/orders [get]
// @Security BearerAuth
func GetAllOrders(c *gin.Context) {
	var orders []models.Order

	err := config.DB.
		Preload("Items.Product").
		Order("created_at DESC").
		Find(&orders).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy vấn đơn hàng"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
