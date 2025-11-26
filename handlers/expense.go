package handlers

import (
	"expense.com/go-backend/models"
	"expense.com/go-backend/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ExpenseHandler contains the handlers for expense operations
type ExpenseHandler struct {
	repo *repository.ExpenseRepository
}

// NewExpenseHandler creates a new instance of ExpenseHandler
func NewExpenseHandler(repo *repository.ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{
		repo: repo,
	}
}

// CreateExpense handles POST /expenses
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var req models.CreateExpenseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	expense, err := h.repo.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Expense created successfully",
		Data:       expense,
	})
}

// GetExpense handles GET /expenses/:id
func (h *ExpenseHandler) GetExpense(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid expense ID",
			Data:       nil,
		})
		return
	}

	expense, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Expense retrieved successfully",
		Data:       expense,
	})
}

// GetAllExpenses handles GET /expenses
func (h *ExpenseHandler) GetAllExpenses(c *gin.Context) {
	expenses := h.repo.GetAll()
	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "All expenses retrieved successfully",
		Data:       expenses,
	})
}

// UpdateExpense handles PUT /expenses/:id
func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid expense ID",
			Data:       nil,
		})
		return
	}

	var req models.UpdateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	expense, err := h.repo.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Expense updated successfully",
		Data:       expense,
	})
}

// DeleteExpense handles DELETE /expenses/:id
func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid expense ID",
			Data:       nil,
		})
		return
	}

	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Expense deleted successfully",
		Data:       nil,
	})
}
