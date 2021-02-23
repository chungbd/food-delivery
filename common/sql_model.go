package common

import "time"

type SQLModel struct {
	ID       int        `json:"id" gorm:"id"`
	Status   int        `json:"status" gorm:"status"`
	CreateAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdateAt *time.Time `json:"updated_at" gorm:"updated_at"`
}
