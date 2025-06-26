// cmd/server/main.go
package main

import (
	"log"
	"net/http"
	"os"

	"debt-tracker-backend/configs"
	"debt-tracker-backend/internal/database"
	"debt-tracker-backend/internal/handlers"
	"debt-tracker-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	config := configs.Load()

	// Initialize database
	db, err := database.Initialize(config.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(config.DatabaseURL); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize Gin router
	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "debt-tracker-api",
		})
	})

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db, config.JWTSecret)
	contactHandler := handlers.NewContactHandler(db)
	debtHandler := handlers.NewDebtHandler(db)
	transactionHandler := handlers.NewTransactionHandler(db)

	// API routes
	api := router.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/me", middleware.AuthRequired(config.JWTSecret), authHandler.GetProfile)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthRequired(config.JWTSecret))
		{
			// Contact routes
			contacts := protected.Group("/contacts")
			{
				contacts.GET("", contactHandler.GetContacts)
				contacts.POST("", contactHandler.CreateContact)
				contacts.GET("/:id", contactHandler.GetContact)
				contacts.PUT("/:id", contactHandler.UpdateContact)
				contacts.DELETE("/:id", contactHandler.DeleteContact)
			}

			// Debt routes
			debts := protected.Group("/debts")
			{
				debts.GET("", debtHandler.GetDebts)
				debts.POST("", debtHandler.CreateDebt)
				debts.GET("/summary", debtHandler.GetDebtSummary)
				debts.GET("/:id", debtHandler.GetDebt)
				debts.PUT("/:id", debtHandler.UpdateDebt)
				debts.DELETE("/:id", debtHandler.DeleteDebt)
				// Debt-specific transactions
				debts.GET("/:id/transactions", transactionHandler.GetDebtTransactions)
			}

			// Transaction routes
			transactions := protected.Group("/transactions")
			{
				transactions.GET("", transactionHandler.GetTransactions)
				transactions.POST("", transactionHandler.CreateTransaction)
				transactions.GET("/:id", transactionHandler.GetTransaction)
				transactions.DELETE("/:id", transactionHandler.DeleteTransaction)
			}
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}