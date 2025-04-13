package service

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type ListItemRepo interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoIem, error)
}

type listItemService struct {
	repo      ListItemRepo
	requester common.Requester
}

func NewListItemService(repo ListItemRepo, requester common.Requester) *listItemService {
	return &listItemService{repo: repo, requester: requester}
}

func (listItemService *listItemService) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoIem, error) {

	data, err := listItemService.repo.ListItem(ctx, filter, paging, moreKeys...)
	if err != nil {
		return nil, err
	}

	return data, nil
}
