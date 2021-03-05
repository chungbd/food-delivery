package notebusiness

import (
	"context"
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
)

type CreateNoteStore interface {
	Create(ctx context.Context, note *notemodel.CreateNote) error
}

type CreateNoteBiz struct {
	store CreateNoteStore
}

func NewCreateNoteBiz(store CreateNoteStore) *CreateNoteBiz {
	return &CreateNoteBiz{store: store}
}

func (biz *CreateNoteBiz) CreateNote(ctx context.Context, data *notemodel.CreateNote) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(data.TableName(), err)
	}

	return nil
}
