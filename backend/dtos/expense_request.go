package dtos

import "time"

type ExpenseRequest struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	UpdateAt    time.Time `json:"update_at"`
	Date		string    `json:"date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
}
