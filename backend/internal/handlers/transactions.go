// internal/handlers/transactions.go
package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"debt-tracker-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	db *sql.DB
}

func NewTransactionHandler(db *sql.DB) *TransactionHandler {
	return &TransactionHandler{db: db}
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	userID := c.GetInt("user_id")

	rows, err := h.db.Query(`
		SELECT t.id, t.debt_id, t.amount, t.transaction_type, t.description, t.created_at
		FROM transactions t
		JOIN debts d ON t.debt_id = d.id
		WHERE d.user_id = ?
		ORDER BY t.created_at DESC
	`, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.ID, &transaction.DebtID, &transaction.Amount,
			&transaction.TransactionType, &transaction.Description, &transaction.CreatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan transaction"})
			return
		}
		transactions = append(transactions, transaction)
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	userID := c.GetInt("user_id")
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	// Check if transaction belongs to user's debt
	var debtExists int
	err = h.db.QueryRow(`
		SELECT d.id FROM transactions t
		JOIN debts d ON t.debt_id = d.id
		WHERE t.id = ? AND d.user_id = ?
	`, transactionID, userID).Scan(&debtExists)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	_, err = h.db.Exec("DELETE FROM transactions WHERE id = ?", transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	userID := c.GetInt("user_id")

	var req models.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if debt belongs to user
	var debtExists int
	err := h.db.QueryRow(
		"SELECT id FROM debts WHERE id = ? AND user_id = ?",
		req.DebtID, userID,
	).Scan(&debtExists)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debt not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	result, err := h.db.Exec(`
		INSERT INTO transactions (debt_id, amount, transaction_type, description) 
		VALUES (?, ?, ?, ?)
	`, req.DebtID, req.Amount, req.TransactionType, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	transactionID, _ := result.LastInsertId()

	var transaction models.Transaction
	err = h.db.QueryRow(`
		SELECT id, debt_id, amount, transaction_type, description, created_at 
		FROM transactions WHERE id = ?
	`, transactionID).Scan(
		&transaction.ID, &transaction.DebtID, &transaction.Amount,
		&transaction.TransactionType, &transaction.Description, &transaction.CreatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get created transaction"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	userID := c.GetInt("user_id")
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	err = h.db.QueryRow(`
		SELECT t.id, t.debt_id, t.amount, t.transaction_type, t.description, t.created_at
		FROM transactions t
		JOIN debts d ON t.debt_id = d.id
		WHERE t.id = ? AND d.user_id = ?
	`, transactionID, userID).Scan(
		&transaction.ID, &transaction.DebtID, &transaction.Amount,
		&transaction.TransactionType, &transaction.Description, &transaction.CreatedAt,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) GetDebtTransactions(c *gin.Context) {
	userID := c.GetInt("user_id")
	debtID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid debt ID"})
		return
	}

	// Check if debt belongs to user
	var debtExists int
	err = h.db.QueryRow(
		"SELECT id FROM debts WHERE id = ? AND user_id = ?",
		debtID, userID,
	).Scan(&debtExists)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debt not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	rows, err := h.db.Query(`
		SELECT id, debt_id, amount, transaction_type, description, created_at
		FROM transactions 
		WHERE debt_id = ?
		ORDER BY created_at DESC
	`, debtID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.ID, &transaction.DebtID, &transaction.Amount,
			&transaction.TransactionType, &transaction.Description, &transaction.CreatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan transaction"})
			return
		}
		transactions = append(transactions, transaction)
	}

	c.JSON(http.StatusOK, transactions)
}