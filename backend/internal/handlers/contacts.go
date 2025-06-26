// internal/handlers/contacts.go
package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"debt-tracker-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	db *sql.DB
}

func NewContactHandler(db *sql.DB) *ContactHandler {
	return &ContactHandler{db: db}
}

func (h *ContactHandler) GetContacts(c *gin.Context) {
	userID := c.GetInt("user_id")

	rows, err := h.db.Query(`
		SELECT id, user_id, name, phone, email, is_active, created_at, updated_at 
		FROM contacts 
		WHERE user_id = ? AND is_active = 1
		ORDER BY name ASC
	`, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get contacts"})
		return
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		err := rows.Scan(
			&contact.ID, &contact.UserID, &contact.Name, &contact.Phone,
			&contact.Email, &contact.IsActive, &contact.CreatedAt, &contact.UpdatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan contact"})
			return
		}
		contacts = append(contacts, contact)
	}

	c.JSON(http.StatusOK, contacts)
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	userID := c.GetInt("user_id")

	var req models.CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.db.Exec(`
		INSERT INTO contacts (user_id, name, phone, email) 
		VALUES (?, ?, ?, ?)
	`, userID, req.Name, req.Phone, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contact"})
		return
	}

	contactID, _ := result.LastInsertId()

	var contact models.Contact
	err = h.db.QueryRow(`
		SELECT id, user_id, name, phone, email, is_active, created_at, updated_at 
		FROM contacts WHERE id = ?
	`, contactID).Scan(
		&contact.ID, &contact.UserID, &contact.Name, &contact.Phone,
		&contact.Email, &contact.IsActive, &contact.CreatedAt, &contact.UpdatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get created contact"})
		return
	}

	c.JSON(http.StatusCreated, contact)
}

func (h *ContactHandler) GetContact(c *gin.Context) {
	userID := c.GetInt("user_id")
	contactID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	var contact models.Contact
	err = h.db.QueryRow(`
		SELECT id, user_id, name, phone, email, is_active, created_at, updated_at 
		FROM contacts WHERE id = ? AND user_id = ?
	`, contactID, userID).Scan(
		&contact.ID, &contact.UserID, &contact.Name, &contact.Phone,
		&contact.Email, &contact.IsActive, &contact.CreatedAt, &contact.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get contact"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) UpdateContact(c *gin.Context) {
	userID := c.GetInt("user_id")
	contactID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	var req models.UpdateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.db.Exec(`
		UPDATE contacts 
		SET name = ?, phone = ?, email = ?, updated_at = CURRENT_TIMESTAMP 
		WHERE id = ? AND user_id = ?
	`, req.Name, req.Phone, req.Email, contactID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update contact"})
		return
	}

	var contact models.Contact
	err = h.db.QueryRow(`
		SELECT id, user_id, name, phone, email, is_active, created_at, updated_at 
		FROM contacts WHERE id = ? AND user_id = ?
	`, contactID, userID).Scan(
		&contact.ID, &contact.UserID, &contact.Name, &contact.Phone,
		&contact.Email, &contact.IsActive, &contact.CreatedAt, &contact.UpdatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get updated contact"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) DeleteContact(c *gin.Context) {
	userID := c.GetInt("user_id")
	contactID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	_, err = h.db.Exec(`
		UPDATE contacts 
		SET is_active = 0, updated_at = CURRENT_TIMESTAMP 
		WHERE id = ? AND user_id = ?
	`, contactID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}