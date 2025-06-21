package controllers

import (

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/config"
	"github.com/sdafbja/telecom-shop/models"
)
// GetAdminStats godoc
// @Summary Thống kê tổng quan cho Admin
// @Description Lấy tổng số người dùng, đơn hàng và tổng doanh thu
// @Tags Admin
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /admin/stats [get]
func GetAdminStats(c *gin.Context) {
	var userCount int64
	var orderCount int64
	var totalRevenue float64

	config.DB.Model(&models.User{}).Count(&userCount)
	config.DB.Model(&models.Order{}).Count(&orderCount)
	config.DB.Model(&models.Order{}).Select("SUM(total)").Scan(&totalRevenue)

	c.JSON(http.StatusOK, gin.H{
		"users":   userCount,
		"orders":  orderCount,
		"revenue": totalRevenue,
	})
}
