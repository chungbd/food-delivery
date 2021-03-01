package notestorage

import (
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
)

func (s *store) ListDataWithCondition(condition map[string]interface{}, paging *common.Paging) ([]notemodel.Note, error) {
	db := s.db

	var data []notemodel.Note

	if err := db.Table(notemodel.Note{}.TableName()).Limit(paging.Limit).
		Offset(paging.Page - 1).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
