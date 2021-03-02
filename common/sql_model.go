package common

import "time"

type SQLModel struct {
	ID        int        `json:"id" gorm:"id"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at" `
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}
