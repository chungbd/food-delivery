package notestorage

import (
	"context"
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
)

func (s *store) ListDataWithCondition(ctx context.Context, condition map[string]interface{}, paging *common.Paging) ([]notemodel.Note, error) {
	db := s.db.Table(notemodel.Note{}.TableName())

	db = db.Where("status <> 0")
	db = db.Where(condition)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	var data []notemodel.Note

	if err := db.Limit(paging.Limit).
		Offset(paging.Page - 1).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
