package notestorage

import (
	"context"
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
)

func (s *store) Update(ctx context.Context, note notemodel.UpdateNote) error {
	db := s.db.Begin()

	db = db.Table(note.TableName())
	db = db.Where("id = ?", note.Id)

	if err := db.Updates(note).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
