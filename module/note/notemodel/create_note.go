package notemodel

import "food-delivery/common"

type CreateNote struct {
	common.SQLModel `json:",inline"`
	UserId          int            `json:"-" gorm:"column:user_id;"`
	Title           string         `json:"title" form:"title" gorm:"column:title" binding:"required"`
	Content         string         `json:"content" form:"content" gorm:"column:content" binding:"required"`
	Cover           *common.Image  `json:"cover"`
	Photos          *common.Images `json:"photos"`
}

// TableName is the table name of note from DB
func (CreateNote) TableName() string {
	return NoteTableName
}
