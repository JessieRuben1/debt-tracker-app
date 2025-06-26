package models

import "time"

type User struct {
	ID           int       `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Name         string    `json:"name" db:"name"`
	Phone        *string   `json:"phone" db:"phone"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type Contact struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Phone     *string   `json:"phone" db:"phone"`
	Email     *string   `json:"email" db:"email"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateContactRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type UpdateContactRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Debt struct {
	ID          int      `json:"id" db:"id"`
	UserID      int      `json:"user_id" db:"user_id"`
	ContactID   int      `json:"contact_id" db:"contact_id"`
	Amount      float64  `json:"amount" db:"amount"`
	Direction   string   `json:"direction" db:"direction"` // "owe_to" or "owe_from"
	Status      string   `json:"status" db:"status"`       // "active", "settled", "removed"
	Description *string  `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Contact     *Contact `json:"contact,omitempty"`
}

type CreateDebtRequest struct {
	ContactID   int     `json:"contact_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Direction   string  `json:"direction" binding:"required,oneof=owe_to owe_from"`
	Description string  `json:"description"`
}

type UpdateDebtRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Status      string  `json:"status" binding:"oneof=active settled removed"`
}

type DebtSummary struct {
	TotalOwedToOthers     float64 `json:"total_owed_to_others"`
	TotalOwedFromOthers   float64 `json:"total_owed_from_others"`
	NetBalance            float64 `json:"net_balance"`
	ActiveDebtsCount      int     `json:"active_debts_count"`
	ContactsWithDebts     int     `json:"contacts_with_debts"`
}

type Transaction struct {
	ID              int       `json:"id" db:"id"`
	DebtID          int       `json:"debt_id" db:"debt_id"`
	Amount          float64   `json:"amount" db:"amount"`
	TransactionType string    `json:"transaction_type" db:"transaction_type"`
	Description     *string   `json:"description" db:"description"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

type CreateTransactionRequest struct {
	DebtID          int     `json:"debt_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
	TransactionType string  `json:"transaction_type" binding:"required,oneof=lent borrowed paid_back received_back"`
	Description     string  `json:"description"`
}
