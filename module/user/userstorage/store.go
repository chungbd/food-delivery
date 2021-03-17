package userstorage

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

// NewSQLStore - Creating store pointer
func NewSQLStore(db *gorm.DB) *store {
	return &store{db: db}
}
