package notestorage

import (
	"context"
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
	"gorm.io/gorm"
)

func (s *store) FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*notemodel.Note, error) {
	db := s.db

	var note notemodel.Note

	if err := db.Table(note.TableName()).Where(condition).First(&note).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &note, nil
}
