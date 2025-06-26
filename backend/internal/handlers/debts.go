// internal/handlers/debts.go
package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"debt-tracker-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type DebtHandler struct {
	db *sql.DB
}

func NewDebtHandler(db *sql.DB) *DebtHandler {
	return &DebtHandler{db: db}
}

func (h *DebtHandler) GetDebts(c *gin.Context) {
	userID := c.GetInt("user_id")

	rows, err := h.db.Query(`
		SELECT d.id, d.user_id, d.contact_id, d.amount, d.direction, d.status, 
		       d.description, d.created_at, d.updated_at,
		       c.id, c.name, c.phone, c.email
		FROM debts d
		JOIN contacts c ON d.contact_id = c.id
		WHERE d.user_id = ? AND d.status = 'active'
		ORDER BY d.created_at DESC
	`, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get debts"})
		return
	}
	defer rows.Close()

	var debts []models.Debt
	for rows.Next() {
		var debt models.Debt
		var contact models.Contact
		
		err := rows.Scan(
			&debt.ID, &debt.UserID, &debt.ContactID, &debt.Amount, &debt.Direction,
			&debt.Status, &debt.Description, &debt.CreatedAt, &debt.UpdatedAt,
			&contact.ID, &contact.Name, &contact.Phone, &contact.Email,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan debt"})
			return
		}
		
		debt.Contact = &contact
		debts = append(debts, debt)
	}

	c.JSON(http.StatusOK, debts)
}

func (h *DebtHandler) CreateDebt(c *gin.Context) {
	userID := c.GetInt("user_id")

	var req models.CreateDebtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if contact belongs to user
	var contactExists int
	err := h.db.QueryRow(
		"SELECT id FROM contacts WHERE id = ? AND user_id = ? AND is_active = 1",
		req.ContactID, userID,
	).Scan(&contactExists)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contact not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Create the debt (allow multiple debts per contact by removing unique constraint check)
	result, err := h.db.Exec(`
		INSERT INTO debts (user_id, contact_id, amount, direction, description) 
		VALUES (?, ?, ?, ?, ?)
	`, userID, req.ContactID, req.Amount, req.Direction, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create debt"})
		return
	}

	debtID, _ := result.LastInsertId()

	var debt models.Debt
	var contact models.Contact
	err = h.db.QueryRow(`
		SELECT d.id, d.user_id, d.contact_id, d.amount, d.direction, d.status, 
		       d.description, d.created_at, d.updated_at,
		       c.id, c.name, c.phone, c.email
		FROM debts d
		JOIN contacts c ON d.contact_id = c.id
		WHERE d.id = ?
	`, debtID).Scan(
		&debt.ID, &debt.UserID, &debt.ContactID, &debt.Amount, &debt.Direction,
		&debt.Status, &debt.Description, &debt.CreatedAt, &debt.UpdatedAt,
		&contact.ID, &contact.Name, &contact.Phone, &contact.Email,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get created debt"})
		return
	}

	debt.Contact = &contact
	c.JSON(http.StatusCreated, debt)
}

func (h *DebtHandler) GetDebt(c *gin.Context) {
	userID := c.GetInt("user_id")
	debtID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid debt ID"})
		return
	}

	var debt models.Debt
	var contact models.Contact
	err = h.db.QueryRow(`
		SELECT d.id, d.user_id, d.contact_id, d.amount, d.direction, d.status, 
		       d.description, d.created_at, d.updated_at,
		       c.id, c.name, c.phone, c.email
		FROM debts d
		JOIN contacts c ON d.contact_id = c.id
		WHERE d.id = ? AND d.user_id = ?
	`, debtID, userID).Scan(
		&debt.ID, &debt.UserID, &debt.ContactID, &debt.Amount, &debt.Direction,
		&debt.Status, &debt.Description, &debt.CreatedAt, &debt.UpdatedAt,
		&contact.ID, &contact.Name, &contact.Phone, &contact.Email,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Debt not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get debt"})
		return
	}

	debt.Contact = &contact
	c.JSON(http.StatusOK, debt)
}

func (h *DebtHandler) UpdateDebt(c *gin.Context) {
	userID := c.GetInt("user_id")
	debtID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid debt ID"})
		return
	}

	var req models.UpdateDebtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.db.Exec(`
		UPDATE debts 
		SET amount = ?, description = ?, status = ?, updated_at = CURRENT_TIMESTAMP 
		WHERE id = ? AND user_id = ?
	`, req.Amount, req.Description, req.Status, debtID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update debt"})
		return
	}

	var debt models.Debt
	var contact models.Contact
	err = h.db.QueryRow(`
		SELECT d.id, d.user_id, d.contact_id, d.amount, d.direction, d.status, 
		       d.description, d.created_at, d.updated_at,
		       c.id, c.name, c.phone, c.email
		FROM debts d
		JOIN contacts c ON d.contact_id = c.id
		WHERE d.id = ? AND d.user_id = ?
	`, debtID, userID).Scan(
		&debt.ID, &debt.UserID, &debt.ContactID, &debt.Amount, &debt.Direction,
		&debt.Status, &debt.Description, &debt.CreatedAt, &debt.UpdatedAt,
		&contact.ID, &contact.Name, &contact.Phone, &contact.Email,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get updated debt"})
		return
	}

	debt.Contact = &contact
	c.JSON(http.StatusOK, debt)
}

func (h *DebtHandler) DeleteDebt(c *gin.Context) {
	userID := c.GetInt("user_id")
	debtID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid debt ID"})
		return
	}

	_, err = h.db.Exec(`
		UPDATE debts 
		SET status = 'removed', updated_at = CURRENT_TIMESTAMP 
		WHERE id = ? AND user_id = ?
	`, debtID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete debt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Debt deleted successfully"})
}

func (h *DebtHandler) GetDebtSummary(c *gin.Context) {
	userID := c.GetInt("user_id")

	var summary models.DebtSummary

	// Calculate totals
	err := h.db.QueryRow(`
		SELECT 
			COALESCE(SUM(CASE WHEN direction = 'owe_to' AND status = 'active' THEN amount ELSE 0 END), 0) as total_owed_to,
			COALESCE(SUM(CASE WHEN direction = 'owe_from' AND status = 'active' THEN amount ELSE 0 END), 0) as total_owed_from,
			COUNT(CASE WHEN status = 'active' THEN 1 END) as active_count,
			COUNT(DISTINCT CASE WHEN status = 'active' THEN contact_id END) as contacts_count
		FROM debts 
		WHERE user_id = ?
	`, userID).Scan(&summary.TotalOwedToOthers, &summary.TotalOwedFromOthers, 
		&summary.ActiveDebtsCount, &summary.ContactsWithDebts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get debt summary"})
		return
	}

	summary.NetBalance = summary.TotalOwedFromOthers - summary.TotalOwedToOthers

	c.JSON(http.StatusOK, summary)
}