package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"social-todo-list/common"
	"social-todo-list/component/tokenprovider"
	"social-todo-list/modules/user/model"
	"strings"
)

type AuthStore interface {
	FindUser(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	// "Authorization" : "Bearer {token}

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

// Required Auth
// 1. Get token from Header
// 2. Validate token and parse to payload
// 3. From the token payload, we ues user_id to find from DB

func RequiredAuth(authStore AuthStore, tokenProvider tokenprovider.Provider) func(c *gin.Context) {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.Request.Header.Get("Authorization"))

		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId()})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
