package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yourusername/ecommerce-platform/backend/handlers"
	"github.com/yourusername/ecommerce-platform/backend/middlewares"
	"github.com/yourusername/ecommerce-platform/backend/models"
)

func main() {
	r := gin.Default()

	db, err := gorm.Open("postgres", "host=localhost dbname=ecommerce sslmode=disable user=youruser password=yourpassword")
	if err != nil {
		panic("Failed to connect to database.")
	}
	defer db.Close()

	db.AutoMigrate(&models.Product{})

	// Seed the database with sample data
	db.Create(&models.Product{Name: "Product 1", Description: "Sample product 1", Price: 19.99})
	db.Create(&models.Product{Name: "Product 2", Description: "Sample product 2", Price: 29.99})

	authHandler := handlers.NewAuthHandler(db)
	productHandler := handlers.NewProductHandler(db)
	paymentHandler := handlers.NewPaymentHandler(db)

	jwtAuthMiddleware := middlewares.JWTAuthMiddleware("your_secret_key")

	api := r.Group("/api")
	api.Use(jwtAuthMiddleware)

	r.POST("/login", authHandler.Login)

	api.GET("/products", productHandler.GetProducts)
	api.POST("/payments/stripe", paymentHandler.CreateStripeCharge)
	api.POST("/payments/paypal", paymentHandler.CreatePaypalPayment)
	api.GET("/payments/ethereum/:txHash", paymentHandler.VerifyEthereumTransaction)

	// Implement other API endpoints here.

	r.Run(":8080")
}
