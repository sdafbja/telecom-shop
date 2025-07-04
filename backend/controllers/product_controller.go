package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"path/filepath"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/config"
	"github.com/sdafbja/telecom-shop/models"
)

// Lấy danh sách sản phẩm
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	categoryID := c.Query("category_id")
	search := strings.ToLower(c.Query("search"))

	query := config.DB.Model(&models.Product{}).Where("is_deleted = false").Preload("Category")

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if search != "" {
		query = query.Where(
			"LOWER(name) LIKE ? OR LOWER(description) LIKE ?",
			"%"+search+"%", "%"+search+"%",
		)
	}

	if err := query.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách sản phẩm"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// Tạo sản phẩm mới
func CreateProduct(c *gin.Context) {
	var product models.Product

	name := c.PostForm("name")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")
	categoryIDStr := c.PostForm("category_id")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Giá không hợp lệ"})
		return
	}

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Danh mục không hợp lệ"})
		return
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.CategoryID = uint(categoryID)

	file, err := c.FormFile("image")
	if err == nil {
		filename := filepath.Base(file.Filename)
		savePath := "uploads/" + filename
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu ảnh"})
			return
		}
		product.ImageURL = "/images/" + filename
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo sản phẩm"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// ❗️Xoá mềm sản phẩm
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sản phẩm"})
		return
	}

	if err := config.DB.Model(&product).Update("is_deleted", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xoá sản phẩm"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sản phẩm đã được đánh dấu là đã xoá"})
}

// Lấy chi tiết sản phẩm
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("🔍 Gọi API chi tiết sản phẩm ID =", id)

	var product models.Product
	if err := config.DB.First(&product, "id = ? AND is_deleted = false", id).Error; err != nil {
		fmt.Println("❌ Không tìm thấy sản phẩm:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sản phẩm"})
		return
	}

	fmt.Println("✅ Tìm thấy sản phẩm:", product.Name)
	c.JSON(http.StatusOK, product)
}

// Cập nhật sản phẩm
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, "id = ? AND is_deleted = false", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sản phẩm"})
		return
	}

	if name := c.PostForm("name"); name != "" {
		product.Name = name
	}
	if desc := c.PostForm("description"); desc != "" {
		product.Description = desc
	}
	if priceStr := c.PostForm("price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			product.Price = price
		}
	}
	if categoryIDStr := c.PostForm("category_id"); categoryIDStr != "" {
		if categoryID, err := strconv.Atoi(categoryIDStr); err == nil {
			product.CategoryID = uint(categoryID)
		}
	}

	file, err := c.FormFile("image")
	if err == nil {
		filename := filepath.Base(file.Filename)
		savePath := "uploads/" + filename
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu ảnh"})
			return
		}
		product.ImageURL = "/images/" + filename
	}

	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật sản phẩm"})
		return
	}

	c.JSON(http.StatusOK, product)
}
