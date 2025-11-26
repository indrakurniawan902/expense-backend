package main

import (
	"fmt"
	"os"

	"expense.com/go-backend/handlers"
	"expense.com/go-backend/models"
	"expense.com/go-backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize repository
	expenseRepo := repository.NewExpenseRepository()

	// Initialize handler
	expenseHandler := handlers.NewExpenseHandler(expenseRepo)

	// Setup Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(corsMiddleware())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.APIResponse{
			StatusCode: http.StatusOK,
			Message:    "Expense API is running",
			Data:       gin.H{"status": "ok"},
		})
	})

	// Expense routes
	expenses := router.Group("/expenses")
	{
		expenses.POST("", expenseHandler.CreateExpense)           // Create expense
		expenses.GET("", expenseHandler.GetAllExpenses)           // List all expenses
		expenses.GET("/:id", expenseHandler.GetExpense)           // Get single expense
		expenses.PUT("/:id", expenseHandler.UpdateExpense)        // Update expense
		expenses.DELETE("/:id", expenseHandler.DeleteExpense)     // Delete expense
	}

	// Start server
	port := getPort()
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("🚀 Expense API Server starting on %s\n", addr)
	fmt.Println("✅ API is ready to accept requests")
	router.Run(addr)
}

// getPort returns the port from environment or default to 8080
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

// corsMiddleware adds CORS headers
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
