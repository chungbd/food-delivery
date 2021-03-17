package notestorage

import (
	"context"
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
)

func (s *store) Create(ctx context.Context, note *notemodel.CreateNote) error {
	db := s.db.Begin()
	db = db.Table(note.TableName())
	note.Status = 1
	if err := db.Create(note).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
