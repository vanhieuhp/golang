package service

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/component/tokenprovider"
	"social-todo-list/modules/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
}

type LoginService struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginService(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, expire int) *LoginService {
	return &LoginService{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expire,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1 Access token and refresh token
// 4. Return token(s)

func (service *LoginService) Login(context context.Context, data *model.UserLogin) (tokenprovider.Token, error) {
	user, err := service.storeUser.FindUser(context, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	passHashed := service.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	payload := &common.TokenPayload{
		UId:   user.Id,
		URole: user.Role.String(),
	}

	accessToken, err := service.tokenProvider.Generate(payload, service.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
