package notebusiness

import (
	"context"
	"errors"
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
)

type UpdateNoteStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*notemodel.Note, error)
	Update(ctx context.Context, note notemodel.UpdateNote) error
}

type updateNoteBiz struct {
	store UpdateNoteStore
}

func NewUpdateNoteBiz(store UpdateNoteStore) *updateNoteBiz {
	return &updateNoteBiz{store: store}
}

func (biz *updateNoteBiz) UpdateNote(ctx context.Context, data notemodel.UpdateNote) error {
	// Find note by id
	// if old data has status is 0
	// => error: note has been deleted
	// => error: note have been deleted
	// else
	// delete note
	note, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": data.Id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound(notemodel.EntityName, err)
		}
		return err
	}

	if note.Status == 0 {
		//return common.ErrCannotDeleteEntity()
		return common.ErrDB(errors.New("note has been deleted before"))
	}

	if err := biz.store.Update(ctx, data); err != nil {
		return common.ErrCannotUpdateEntity(note.TableName(), err)
	}

	return nil
}
