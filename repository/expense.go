package repository

import (
	"errors"
	"expense.com/go-backend/models"
	"sync"
	"time"
)

// ExpenseRepository handles all database operations for expenses
type ExpenseRepository struct {
	expenses map[int]*models.Expense
	mu       sync.RWMutex
	nextID   int
}

// NewExpenseRepository creates a new instance of ExpenseRepository
func NewExpenseRepository() *ExpenseRepository {
	return &ExpenseRepository{
		expenses: make(map[int]*models.Expense),
		nextID:   1,
	}
}

// Create adds a new expense to the repository
func (r *ExpenseRepository) Create(req *models.CreateExpenseRequest) (*models.Expense, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("invalid date format, use YYYY-MM-DD")
	}

	expense := &models.Expense{
		ID:          r.nextID,
		Description: req.Description,
		Amount:      req.Amount,
		Category:    req.Category,
		Date:        date,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	r.expenses[r.nextID] = expense
	r.nextID++

	return expense, nil
}

// GetByID retrieves an expense by its ID
func (r *ExpenseRepository) GetByID(id int) (*models.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	expense, exists := r.expenses[id]
	if !exists {
		return nil, errors.New("expense not found")
	}

	return expense, nil
}

// GetAll retrieves all expenses
func (r *ExpenseRepository) GetAll() []*models.Expense {
	r.mu.RLock()
	defer r.mu.RUnlock()

	expenses := make([]*models.Expense, 0, len(r.expenses))
	for _, expense := range r.expenses {
		expenses = append(expenses, expense)
	}

	return expenses
}

// Update updates an existing expense
func (r *ExpenseRepository) Update(id int, req *models.UpdateExpenseRequest) (*models.Expense, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	expense, exists := r.expenses[id]
	if !exists {
		return nil, errors.New("expense not found")
	}

	if req.Description != "" {
		expense.Description = req.Description
	}
	if req.Amount > 0 {
		expense.Amount = req.Amount
	}
	if req.Category != "" {
		expense.Category = req.Category
	}
	if req.Date != "" {
		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			return nil, errors.New("invalid date format, use YYYY-MM-DD")
		}
		expense.Date = date
	}

	expense.UpdatedAt = time.Now()
	r.expenses[id] = expense

	return expense, nil
}

// Delete removes an expense from the repository
func (r *ExpenseRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.expenses[id]; !exists {
		return errors.New("expense not found")
	}

	delete(r.expenses, id)
	return nil
}
