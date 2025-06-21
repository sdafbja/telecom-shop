package controllers


import (
	"net/http"
	"strconv"
	"strings"

	// "os"
	"path/filepath"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/config"
	"github.com/sdafbja/telecom-shop/models"
)
// GetAllProducts godoc
// @Summary Lấy danh sách sản phẩm
// @Description Lấy tất cả sản phẩm, có thể lọc theo category_id và tìm kiếm theo tên
// @Tags Products
// @Produce json
// @Param category_id query string false "ID danh mục"
// @Param search query string false "Từ khóa tìm kiếm"
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
// Lấy tất cả sản phẩm (có thể lọc theo category_id hoặc tìm kiếm theo tên)
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	categoryID := c.Query("category_id")
	search := strings.ToLower(c.Query("search"))

	query := config.DB.Model(&models.Product{}).Preload("Category")

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if search != "" {
		query = query.Where(
			"LOWER(products.name) LIKE ? OR LOWER(products.description) LIKE ?",
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
// CreateProduct godoc
// @Summary Tạo sản phẩm mới
// @Description Tạo sản phẩm với thông tin và ảnh tải lên (chỉ Admin)
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Tên sản phẩm"
// @Param description formData string false "Mô tả"
// @Param price formData number true "Giá"
// @Param category_id formData int true "ID danh mục"
// @Param image formData file false "Ảnh sản phẩm"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
// @Security BearerAuth
func CreateProduct(c *gin.Context) {
	var product models.Product

	// Lấy dữ liệu từ multipart form
	name := c.PostForm("name")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")
	categoryIDStr := c.PostForm("category_id")

	//  Ép kiểu dữ liệu
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

	// Gán dữ liệu
	product.Name = name
	product.Description = description
	product.Price = price
	product.CategoryID = uint(categoryID)

	
	file, err := c.FormFile("image")
	if err == nil {
		filename := filepath.Base(file.Filename)
		savePath := "uploads/" + filename // nơi lưu thật
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu ảnh"})
			return
		}
		product.ImageURL = "/images/" + filename // ✅ đường dẫn public trả về
	}

	// ✅ Tạo sản phẩm trong DB
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo sản phẩm"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// Xoá sản phẩm
// DeleteProduct godoc
// @Summary Xoá sản phẩm
// @Description Xoá sản phẩm theo ID (chỉ Admin)
// @Tags Products
// @Produce json
// @Param id path int true "ID sản phẩm"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [delete]
// @Security BearerAuth
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	if err := config.DB.Delete(&models.Product{}, productID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xoá sản phẩm"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đã xoá sản phẩm"})
}

// Lấy chi tiết sản phẩm theo ID
// GetProductByID godoc
// @Summary Lấy chi tiết sản phẩm
// @Description Lấy thông tin sản phẩm theo ID
// @Tags Products
// @Produce json
// @Param id path int true "ID sản phẩm"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("🔍 Gọi API chi tiết sản phẩm ID =", id)

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		fmt.Println("❌ Không tìm thấy sản phẩm:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sản phẩm"})
		return
	}

	fmt.Println("✅ Tìm thấy sản phẩm:", product.Name)
	c.JSON(http.StatusOK, product)
}
