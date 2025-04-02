package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"social-todo-list/common"
	"social-todo-list/modules/user/model"
)

func (sql *sqlStorage) FindUser(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error) {
	m := model.User{}
	db := sql.db.Table(m.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user model.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrDB(err)
		}
		return nil, err
	}

	return &user, nil
}