package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/config"
	"github.com/sdafbja/telecom-shop/models"
	"github.com/sdafbja/telecom-shop/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
	"strings"
)

// Register godoc
// @Summary Đăng ký tài khoản
// @Description Đăng ký người dùng mới với email, mật khẩu và tên
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "Thông tin đăng ký"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
// ✅ Đăng ký
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Gán role mặc định nếu không có (user)
	if input.Role == "" {
		input.Role = "user"
	}

	// Mã hóa mật khẩu
	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	input.Password = string(hashed)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo người dùng"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đăng ký thành công"})
}

// Login godoc
// @Summary Đăng nhập
// @Description Xác thực người dùng và trả về token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "Thông tin đăng nhập"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
// ✅ Đăng nhập
func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Sai email hoặc chưa đăng ký"})
		return
	}

	// So sánh mật khẩu
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Sai mật khẩu"})
		return
	}

	// Tạo JWT token
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo token"})
		return
	}

	// ✅ Trả về token và thông tin người dùng (ẩn mật khẩu)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
func Logout(c *gin.Context) {
	token := extractToken(c)
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token không hợp lệ"})
		return
	}

	// Lưu token vào Redis blacklist với thời gian sống là 3 giờ
	err := config.RedisClient.Set(config.Ctx, "blacklist:"+token, "true", 3*time.Hour).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout thành công!"})
}
func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return ""
}


