package models

// User represents a participant in the expense tracker
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Split represents how much a specific user owes for an expense
type Split struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

// Expense represents a shared cost
type Expense struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	TotalAmount float64 `json:"total_amount"`
	PaidBy      string  `json:"paid_by"` // User ID of who paid
	Splits      []Split `json:"splits"`  // Breakdown of who owes what
}

// Settlement represents a transaction suggested to clear debts
type Settlement struct {
	From   string  `json:"from"`   // User who pays
	To     string  `json:"to"`     // User who receives
	Amount float64 `json:"amount"`
}

// UserBalance represents the net position of a user
type UserBalance struct {
	UserID string  `json:"user_id"`
	Name   string  `json:"name"`
	Amount float64 `json:"balance"` // Positive means they are owed, negative means they owe
}
