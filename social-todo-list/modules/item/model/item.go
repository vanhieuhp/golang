package model

import (
	"errors"
	"social-todo-list/common"
)

const (
	EntityName = "item"
)

var (
	ErrTitleIsBlank  = errors.New("title is blank")
	ErrItemIdDeleted = errors.New("item is deleted")
	ErrItemDeletedNew = common.NewCustomError(errors.New("item is deleted"), "item has been deleted", "ErrItemDeleted")
)

type TodoIem struct {
	common.SQLModel
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoIem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:titles"`
	Description string      `json:"description" gorm:"column:description"`
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
