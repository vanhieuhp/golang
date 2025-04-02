package service

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/user/model"
)

type RegisterStorage interface {
	FindUser(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
	CreateUser(ctx context.Context, data *model.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerService struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterService(registerStorage RegisterStorage, hasher Hasher) *registerService {
	return &registerService{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (business *registerService) Register(context context.Context, data *model.UserCreate) error {
	user, _ := business.registerStorage.FindUser(context, map[string]interface{}{"email": data.Email})

	if user != nil {
		return model.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" //  hard code

	if err := business.registerStorage.CreateUser(context, data); err != nil {
		return common.ErrCannotCreateEntity(model.User{}.TableName(), err)
	}

	return nil
}
