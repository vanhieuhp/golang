package model

import (
	"errors"
	"social-todo-list/common"
)

const (
	EntityName = "items"
)

var (
	ErrTitleIsBlank   = errors.New("title is blank")
	ErrItemIdDeleted  = errors.New("item is deleted")
	ErrItemDeletedNew = common.NewCustomError(errors.New("item is deleted"), "item has been deleted", "ErrItemDeleted")
)

type TodoIem struct {
	common.SQLModel
	Title       string             `json:"title" gorm:"column:title"`
	Description string             `json:"description" gorm:"column:description"`
	Status      *ItemStatus        `json:"status" gorm:"column:status"`
	UserId      int                `json:"user_id" gorm:"column:user_id"`
	Owner       *common.SimpleUser `json:"owner" gorm:"foreignKey:UserId;"`
	LikedCount  int                `json:"liked_count" gorm:"-"`
}

func (TodoIem) TableName() string {
	return EntityName
}

func (i *TodoIem) Mask() {
	i.SQLModel.Mask(common.DbTypeItem)

	if v := i.Owner; i != nil {
		v.Mask()
	}
}

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	UserId      int         `json:"-" gorm:"column:user_id"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItemUpdate) TableName() string {
	return TodoIem{}.TableName()
}

func (TodoItemCreation) TableName() string {
	return TodoIem{}.TableName()
}
