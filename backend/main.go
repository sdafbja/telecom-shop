package main

import (
	"github.com/gin-contrib/cors"
	_ "github.com/sdafbja/telecom-shop/docs"
	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/config"
	"github.com/sdafbja/telecom-shop/models"
	"github.com/sdafbja/telecom-shop/middlewares"
	"github.com/sdafbja/telecom-shop/routes"

	"log"

	"github.com/swaggo/files"           // swagger embed files
	"github.com/swaggo/gin-swagger"     // gin-swagger middleware
)

// @title        Telecom Shop API
// @version      1.0
// @description  API tài liệu cho hệ thống bán hàng viễn thông
// @contact.name Nguyễn Văn Phong
// @contact.email vanphong160514@gmail.com
// @host         localhost:8080
// @BasePath     /
// @schemes      http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Kết nối DB & Redis
	config.ConnectDB()
	config.ConnectRedis()

	// Tự động migrate
	config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Category{},
		&models.Cart{},
		&models.Order{},
		&models.OrderItem{},
		&models.Review{},
	)

	r := gin.Default()


	// Middleware
	r.Use(middlewares.LoggerMiddleware())

	// Static file (ảnh)
	r.Static("/images", "./public/images")

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CORS cấu hình
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Đăng ký các route
	routes.RegisterRoutes(r)

	// Chạy server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Không thể chạy server:", err)
	}
}
