package notebusiness

import (
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
)

type ListNoteStore interface {
	ListDataWithCondition(condition map[string]interface{}, paging *common.Paging) ([]notemodel.Note, error)
}

type listNoteBiz struct {
	store ListNoteStore
}

func NewListNoteBiz(store ListNoteStore) *listNoteBiz {
	return &listNoteBiz{store: store}
}

func (biz *listNoteBiz) ListNote(paging *common.Paging) ([]notemodel.Note, error) {
	// List all notes by paging
	// if list is empty
	// => error: list has been empty
	// else
	// list notes
	notes, err := biz.store.ListDataWithCondition(nil, paging)

	if err != nil {
		return nil, err
	}

	return notes, err
}
