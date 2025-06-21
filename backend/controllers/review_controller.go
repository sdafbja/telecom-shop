package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/config"
	"github.com/sdafbja/telecom-shop/models"
)

// Tạo đánh giá
// CreateReview godoc
// @Summary Gửi đánh giá sản phẩm
// @Description Người dùng gửi đánh giá (rating và comment) cho sản phẩm
// @Tags Reviews
// @Accept json
// @Produce json
// @Param review body models.Review true "Thông tin đánh giá"
// @Success 200 {object} models.Review
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews [post]
// @Security BearerAuth
func CreateReview(c *gin.Context) {
	var input models.Review
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	userID := c.GetUint("user_id")
	input.UserID = userID

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu đánh giá"})
		return
	}

	c.JSON(http.StatusOK, input)
}

// Lấy đánh giá theo sản phẩm
// GetReviewsByProductID godoc
// @Summary Lấy đánh giá sản phẩm
// @Description Trả về danh sách đánh giá của một sản phẩm theo ID
// @Tags Reviews
// @Produce json
// @Param id path int true "ID sản phẩm"
// @Success 200 {array} models.Review
// @Failure 500 {object} map[string]string
// @Router /reviews/{id} [get]
func GetReviewsByProductID(c *gin.Context) {
	productID := c.Param("id")
	var reviews []models.Review

	if err := config.DB.Preload("User").Where("product_id = ?", productID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy đánh giá"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
