package repository

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

type ItemLikeStorage interface {
	GetItemLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listItemRepo struct {
	store     ListItemStorage
	likeStore ItemLikeStorage
	requester common.Requester
}

func NewListItemReo(store ListItemStorage, likeStore ItemLikeStorage, requester common.Requester) *listItemRepo {
	return &listItemRepo{store: store, likeStore: likeStore, requester: requester}
}

func (listItemService *listItemRepo) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoIem, error) {

	data, err := listItemService.store.ListItem(ctx, filter, paging, "Owner")
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	ids := make([]int, len(data))

	for i := range ids {
		ids[i] = data[i].Id
	}

	likeUserMap, err := listItemService.likeStore.GetItemLikes(ctx, ids)

	if err != nil {
		return data, nil
	}

	for i := range data {
		data[i].LikedCount = likeUserMap[data[i].Id]
	}

	return data, nil
}
