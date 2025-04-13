package storage

import (
	"context"
	"github.com/akamensky/base58"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/model"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999" // YYYY-MM-DD

func (sql *sqlStorage) ListUsers(
	ctx context.Context,
	itemId int,
	paging *common.Paging) ([]common.SimpleUser, error) {

	var result []model.Like

	db := sql.db.Table(model.Like{}.TableName()).
		Select("item_id").
		Where("item_id = ?", itemId)

	if err := db.Table(model.Like{}.TableName()).Select("user_id").Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := paging.FakeCursor; v != "" {

		decoded, err := base58.Decode(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}

		timeCreated, err := time.Parse(timeLayout, string(decoded))
		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format(timeLayout))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Select("*").
		Order("created_at desc").
		Limit(paging.Limit).
		Preload("User").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i := range users {
		users[i] = *result[i].User
		users[i].UpdatedAt = nil
		users[i].CreatedAt = result[i].CreatedAt
	}

	if len(users) > 0 {
		users[len(result)-1].Mask()
		paging.NextCursor = base58.Encode([]byte(users[len(result)-1].CreatedAt.Format(timeLayout)))
	}

	return users, nil
}

func (sql *sqlStorage) GetItemLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		ItemId int `gorm:"item_id"`
		Count  int `gorm:"count"`
	}

	var listLike []sqlData

	if err := sql.db.Table(model.Like{}.TableName()).Select("item_id, Count(item_id) as count").
		Where("item_id IN (?)", ids).
		Group("item_id").
		Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.ItemId] = item.Count
	}

	return result, nil
}
