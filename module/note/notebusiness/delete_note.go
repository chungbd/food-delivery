package notebusiness

import (
	"errors"
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
)

type DeleteNoteStore interface {
	FindDataWithCondition(condition map[string]interface{}) (*notemodel.Note, error)
	Delete(id int) error
}

type deleteNoteBiz struct {
	store DeleteNoteStore
}

func NewDeleteNoteBiz(store DeleteNoteStore) *deleteNoteBiz {
	return &deleteNoteBiz{store: store}
}

func (biz *deleteNoteBiz) DeleteNote(noteId int) error {
	// Find note by id
	// if old data has status is 0
	// => error: note has been deleted
	// => error: note have been deleted
	// else
	// delete note
	note, err := biz.store.FindDataWithCondition(map[string]interface{}{"id": noteId})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound(notemodel.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(notemodel.EntityName, err)
	}

	if note.Status == 0 {
		return errors.New("note has been deleted before")
	}

	if err := biz.store.Delete(note.ID); err != nil {
		return err
	}

	return nil
}
