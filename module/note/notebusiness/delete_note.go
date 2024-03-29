package notebusiness

import (
	"context"
	"errors"
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
)

type DeleteNoteStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*notemodel.Note, error)
	Delete(ctx context.Context, id int) error
}

type deleteNoteBiz struct {
	store     DeleteNoteStore
	requester common.Requester
}

func NewDeleteNoteBiz(store DeleteNoteStore, requester common.Requester) *deleteNoteBiz {
	return &deleteNoteBiz{store: store, requester: requester}
}

func (biz *deleteNoteBiz) DeleteNote(ctx context.Context, noteId int) error {
	// Find note by id
	// if old data has status is 0
	// => error: note has been deleted
	// => error: note have been deleted
	// else
	// delete note
	note, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": noteId})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound(notemodel.EntityName, err)
		}
		return err
	}

	if note.Status == 0 {
		return common.ErrCannotDeleteEntity(notemodel.NoteTableName, errors.New("note has been deleted before"))
	}

	if note.UserId != biz.requester.GetUserId() {
		return common.ErrNoPermission(errors.New("you are not owner"))
	}

	if err := biz.store.Delete(ctx, note.Id); err != nil {
		return err
	}

	return nil
}
