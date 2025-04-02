package service

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoIem, error)
}

type listItemBiz struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *listItemBiz {
	return &listItemBiz{store: store}
}

func (biz *listItemBiz) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoIem, error) {

	data, err := biz.store.ListItem(ctx, filter, paging, moreKeys...)
	if err != nil {
		return nil, err
	}

	return data, nil
}
