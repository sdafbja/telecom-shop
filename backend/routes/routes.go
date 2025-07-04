package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/controllers"
	"github.com/sdafbja/telecom-shop/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	// Auth
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Product
	productGroup := r.Group("/products")
	{
		productGroup.GET("", controllers.GetAllProducts)
		productGroup.GET("/:id", controllers.GetProductByID)
		productGroup.POST("", middlewares.JWTAuthMiddleware(), middlewares.RequireAdmin(), controllers.CreateProduct)
		productGroup.PUT("/:id", middlewares.JWTAuthMiddleware(), middlewares.RequireAdmin(), controllers.UpdateProduct)
		// productGroup.PUT("/:id", middlewares.JWTAuthMiddleware(), middlewares.RequireAdmin(), controllers.UpdateProduct)
		productGroup.DELETE("/:id", middlewares.JWTAuthMiddleware(), middlewares.RequireAdmin(), controllers.DeleteProduct)
	}

	// User profile
	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		auth.GET("/profile", controllers.GetCurrentUser)
	}

	// Categories
	categoryGroup := r.Group("/categories")
	{
		categoryGroup.GET("", controllers.GetAllCategories)
		categoryGroup.POST("", middlewares.JWTAuthMiddleware(), middlewares.RequireAdmin(), controllers.CreateCategory)
		categoryGroup.PUT("/:id", middlewares.JWTAuthMiddleware(), middlewares.RequireAdmin(), controllers.UpdateCategory)
		categoryGroup.DELETE("/:id", middlewares.JWTAuthMiddleware(), middlewares.RequireAdmin(), controllers.DeleteCategory)
	}

	// Admin
	admin := r.Group("/admin")
	admin.Use(middlewares.JWTAuthMiddleware(), middlewares.RequireAdmin())
	{
		admin.GET("/users", controllers.GetAllUsers)
		admin.GET("/orders", controllers.GetAllOrders)
		admin.GET("/stats", controllers.GetAdminStats)
		admin.PUT("/users/:id", controllers.UpdateUser)
		admin.DELETE("/users/:id", controllers.DeleteUser)
	}

	// Cart
	cart := r.Group("/cart")
	cart.Use(middlewares.JWTAuthMiddleware())
	{
		cart.POST("", controllers.AddToCart)
		cart.GET("", controllers.GetCart)
		cart.DELETE("/:id", controllers.RemoveFromCart)
		cart.PUT("/:id", controllers.UpdateCartQuantity)


	}

	//Orders
	order := r.Group("/orders")
	order.Use(middlewares.JWTAuthMiddleware())
	{
		order.POST("", controllers.CreateOrder)
		order.GET("", controllers.GetMyOrders)
	}

	review := r.Group("/reviews")
	review.Use(middlewares.JWTAuthMiddleware()) 
	{
		review.POST("", controllers.CreateReview)
	}
	r.GET("/reviews/product/:id", controllers.GetReviewsByProductID)


}
