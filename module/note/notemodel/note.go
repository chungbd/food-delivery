package notemodel

import (
	"food-delivery/common"
)

// NoteTableName is for mysql table name
const EntityName = "Note"
const NoteTableName = "notes"

var (
	ErrNoteDeleted = common.NewCustomError(nil, "note has been deleted before", "ErrNoteDeleted")
)

// Note is a entity for note
type Note struct {
	// ID      int    `json:"id" gorm:"column:id"`
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title"`
	Content         string `json:"content" gorm:"column:content"`
	// Status   uint      `json:"status" gorm:"column:status"`
	// CreateAt time.Time `json:"created_at" gorm:"column:created_at"`
	// UpdateAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName is the table name of note from DB
func (Note) TableName() string {
	return NoteTableName
}
