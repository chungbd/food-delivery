package notestorage

import "food-delivery/module/note/notemodel"

func (s *store) FindDataWithCondition(condition map[string]interface{}) (*notemodel.Note, error) {
	db := s.db

	var data notemodel.Note

	if err := db.Table(data.TableName()).Where(condition).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
