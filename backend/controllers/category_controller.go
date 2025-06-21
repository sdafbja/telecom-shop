package controllers

import (
	"net/http"
	"github.com/sdafbja/telecom-shop/config"
	"github.com/sdafbja/telecom-shop/models"

	"github.com/gin-gonic/gin"
)
// CreateCategory godoc
// @Summary Tạo danh mục mới
// @Description Tạo mới một danh mục sản phẩm (chỉ Admin)
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Thông tin danh mục"
// @Success 201 {object} models.Category
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories [post]
// @Security BearerAuth
// Tạo danh mục mới (admin)
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}
	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo danh mục"})
		return
	}
	c.JSON(http.StatusCreated, category)
}
// GetAllCategories godoc
// @Summary Lấy tất cả danh mục
// @Description Lấy danh sách tất cả danh mục sản phẩm
// @Tags Categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
// Lấy danh sách danh mục
func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	config.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

// UpdateCategory godoc
// @Summary Cập nhật danh mục
// @Description Cập nhật thông tin danh mục theo ID (chỉ Admin)
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "ID danh mục"
// @Param category body models.Category true "Dữ liệu cập nhật"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [put]
// @Security BearerAuth
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy danh mục"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	config.DB.Save(&category)
	c.JSON(http.StatusOK, category)
}


// DeleteCategory godoc
// @Summary Xoá danh mục
// @Description Xoá danh mục sản phẩm theo ID (chỉ Admin)
// @Tags Categories
// @Produce json
// @Param id path int true "ID danh mục"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [delete]
// @Security BearerAuth
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xoá danh mục"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Đã xoá danh mục"})
}

