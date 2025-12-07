package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vanhieuhp/social/internal/store"
)

type userKey string

const userCtx userKey = "user"

func (app *application) getUserHandler(writer http.ResponseWriter, request *http.Request) {
	user := getUserContext(request)

	if err := app.jsonResponse(writer, http.StatusOK, user); err != nil {
		app.internalServerResponse(writer, request, err)
	}
}

type FollowUser struct {
	UserId int64 `json:"user_id"`
}

func (app *application) followUserHandler(writer http.ResponseWriter, request *http.Request) {
	followerUser := getUserContext(request)

	var payload FollowUser
	if err := readJSON(writer, request, &payload); err != nil {
		app.badRequestResponse(writer, request, err)
		return
	}

	ctx := request.Context()

	if err := app.store.Followers.Follow(ctx, followerUser.ID, payload.UserId); err != nil {
		switch {
		case errors.Is(err, store.ErrorConflict):
			app.ConflictResponse(writer, request, err)
		default:
			app.internalServerResponse(writer, request, err)
		}
		return
	}

	if err := app.jsonResponse(writer, http.StatusNoContent, nil); err != nil {
		app.internalServerResponse(writer, request, err)
	}
}

func (app *application) unfollowUserHandler(writer http.ResponseWriter, request *http.Request) {
	followerUser := getUserContext(request)

	var payload FollowUser
	if err := readJSON(writer, request, &payload); err != nil {
		app.badRequestResponse(writer, request, err)
		return
	}

	ctx := request.Context()

	if err := app.store.Followers.Unfollow(ctx, followerUser.ID, payload.UserId); err != nil {
		app.internalServerResponse(writer, request, err)
		return
	}

	if err := app.jsonResponse(writer, http.StatusNoContent, nil); err != nil {
		app.internalServerResponse(writer, request, err)
	}
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		idParam := chi.URLParam(request, "userId")
		userId, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.internalServerResponse(writer, request, err)
		}

		ctx := request.Context()

		user, err := app.store.Users.GetById(ctx, userId)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrorNotFound):
				app.badRequestResponse(writer, request, err)
				return
			default:
				app.internalServerResponse(writer, request, err)
				return
			}
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func getUserContext(request *http.Request) *store.User {
	user, _ := request.Context().Value(userCtx).(*store.User)
	return user
}
