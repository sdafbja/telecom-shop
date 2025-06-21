package controllers

import (
	"net/http"
	"errors"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/models"
	"github.com/sdafbja/telecom-shop/config"
)

// AddToCart godoc
// @Summary Thêm sản phẩm vào giỏ hàng
// @Description Thêm sản phẩm và số lượng vào giỏ hàng của người dùng
// @Tags Cart
// @Accept json
// @Produce json
// @Param item body models.CartInput true "Thông tin sản phẩm cần thêm"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart [post]
// @Security BearerAuth
func AddToCart(c *gin.Context) {
	var input models.CartInput

	if err := c.ShouldBindJSON(&input); err != nil || input.ProductID == 0 || input.Quantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ hoặc thiếu product_id"})
		return
	}

	userID := c.GetUint("user_id")

	// Kiểm tra xem sản phẩm có tồn tại không
	var product models.Product
	if err := config.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sản phẩm không tồn tại"})
		return
	}

	var cart models.Cart
	result := config.DB.Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&cart)

	if result.Error == nil {
		cart.Quantity += input.Quantity
		config.DB.Save(&cart)
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newItem := models.Cart{
			UserID:    userID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		config.DB.Create(&newItem)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đã thêm vào giỏ hàng"})
}

// GetCart godoc
// @Summary Xem giỏ hàng
// @Description Trả về danh sách sản phẩm trong giỏ hàng của người dùng
// @Tags Cart
// @Produce json
// @Success 200 {array} models.Cart
// @Router /cart [get]
// @Security BearerAuth
func GetCart(c *gin.Context) {
	userID := c.GetUint("user_id")
	var cart []models.Cart

	config.DB.Preload("Product").Where("user_id = ?", userID).Find(&cart)
	c.JSON(http.StatusOK, cart)
}

// RemoveFromCart godoc
// @Summary Xoá sản phẩm khỏi giỏ hàng
// @Description Xoá 1 mục trong giỏ hàng theo ID
// @Tags Cart
// @Produce json
// @Param id path int true "ID mục giỏ hàng"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/{id} [delete]
// @Security BearerAuth
func RemoveFromCart(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Cart{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xoá sản phẩm khỏi giỏ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đã xoá sản phẩm khỏi giỏ hàng"})
}

// UpdateCartQuantity godoc
// @Summary Cập nhật số lượng sản phẩm trong giỏ hàng
// @Description Cập nhật số lượng mới cho 1 mục giỏ hàng
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path int true "ID mục giỏ hàng"
// @Param quantity body models.UpdateCartQuantityInput true "Số lượng mới"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /cart/{id} [put]
// @Security BearerAuth
func UpdateCartQuantity(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateCartQuantityInput

	if err := c.ShouldBindJSON(&input); err != nil || input.Quantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Số lượng không hợp lệ"})
		return
	}

	var cart models.Cart
	if err := config.DB.First(&cart, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy mục giỏ hàng"})
		return
	}

	cart.Quantity = input.Quantity
	config.DB.Save(&cart)

	c.JSON(http.StatusOK, gin.H{"message": "Đã cập nhật số lượng"})
}
