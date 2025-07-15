package models

import "time"

type Expense struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Date        string    `json:"date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	UserId      int       `json:"user_id"`
	UpdateAt    time.Time `json:"update_at"`
}
