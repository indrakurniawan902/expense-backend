package models

import "time"

// Expense represents a single expense entry
type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateExpenseRequest is the request body for creating an expense
type CreateExpenseRequest struct {
	Description string  `json:"description" binding:"required"`
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Category    string  `json:"category" binding:"required"`
	Date        string  `json:"date" binding:"required"`
}

// UpdateExpenseRequest is the request body for updating an expense
type UpdateExpenseRequest struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount" binding:"gt=0"`
	Category    string  `json:"category"`
	Date        string  `json:"date"`
}

// APIResponse is the standard response wrapper for all API endpoints
type APIResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// ErrorResponse is used when there's an error
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       interface{} `json:"data"`
}
