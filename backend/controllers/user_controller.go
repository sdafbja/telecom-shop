package controllers

import (

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/config"
	"github.com/sdafbja/telecom-shop/models"
)

// Lấy tất cả người dùng (Admin)
// GetAllUsers godoc
// @Summary Lấy danh sách người dùng
// @Description Lấy tất cả người dùng trong hệ thống (chỉ Admin)
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /admin/users [get]
// @Security BearerAuth
func GetAllUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy vấn người dùng"})
		return
	}

	// Ẩn mật khẩu
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, users)
}

// Lấy thông tin người dùng hiện tại (JWT đã decode user_id)
// GetCurrentUser godoc
// @Summary Lấy thông tin người dùng hiện tại
// @Description Trả về thông tin tài khoản đang đăng nhập (theo JWT)
// @Tags Users
// @Produce json
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /users/me [get]
// @Security BearerAuth
func GetCurrentUser(c *gin.Context) {
	userID := c.GetUint("user_id")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy người dùng"})
		return
	}

	user.Password = "" // ẩn mật khẩu
	c.JSON(http.StatusOK, user)
}

// Xoá người dùng (Admin)
// DeleteUser godoc
// @Summary Xoá người dùng
// @Description Xoá một người dùng theo ID (chỉ Admin)
// @Tags Users
// @Produce json
// @Param id path int true "ID người dùng"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/users/{id} [delete]
// @Security BearerAuth
func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	userIDFromToken := c.GetUint("user_id")

	// Chuyển idParam sang uint
	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	// Không cho xoá chính mình
	if userIDFromToken == uint(idUint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Không thể tự xoá chính mình"})
		return
	}

	// Tìm user cần xoá
	var user models.User
	if err := config.DB.First(&user, uint(idUint)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy người dùng"})
		return
	}

	// Tiến hành xoá
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Xoá thất bại"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đã xoá người dùng"})
}
// Cập nhật người dùng (Admin)
// UpdateUser godoc
// @Summary Cập nhật người dùng
// @Description Cập nhật thông tin người dùng theo ID (chỉ Admin)
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID người dùng"
// @Param user body models.User true "Dữ liệu cập nhật"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/users/{id} [put]
// @Security BearerAuth
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy người dùng"})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role // Cho phép cập nhật vai trò

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật người dùng"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đã cập nhật người dùng"})
}

