package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"id"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at" `
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}

func (sqlModel *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(sqlModel.Id), dbType, 128)
	sqlModel.FakeId = &uid
}
