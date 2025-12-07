package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vanhieuhp/social/internal/store"
)

type postKey string

const postCtx postKey = "post"

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=10000"`
	Tags    []string `json:"tags"`
}

func (app *application) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err := Validate.Struct(payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// TODO: Change after auth
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerResponse(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.internalServerResponse(w, r, err)
		return
	}
}

func (app *application) getPostHandler(writer http.ResponseWriter, request *http.Request) {
	post := getPostFromCtx(request)

	comments, err := app.store.Comments.GetCommentsByPostId(request.Context(), post.ID)
	if err != nil {
		app.internalServerResponse(writer, request, err)
		return
	}

	post.Comments = comments

	if err := app.jsonResponse(writer, http.StatusOK, post); err != nil {
		app.internalServerResponse(writer, request, err)
		return
	}
}

func (app *application) deletePostHandler(writer http.ResponseWriter, request *http.Request) {
	idParam := chi.URLParam(request, "postId")
	postId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerResponse(writer, request, err)
	}

	ctx := request.Context()

	if err = app.store.Posts.Delete(ctx, postId); err != nil {
		switch {
		case errors.Is(err, store.ErrorNotFound):
			app.notFoundResponse(writer, request, err)

		default:
			app.internalServerResponse(writer, request, err)
		}
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

type UpdatePostPayload struct {
	Title   *string `json:"title" validate:"omitempty,max=100"`
	Content *string `json:"content" validate:"omitempty,max=10000"`
}

func (app *application) updatePostHandler(writer http.ResponseWriter, request *http.Request) {
	post := getPostFromCtx(request)

	var payload UpdatePostPayload
	if err := readJSON(writer, request, &payload); err != nil {
		app.badRequestResponse(writer, request, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(writer, request, err)
		return
	}

	log.Printf("%v", payload)
	if payload.Content != nil {
		post.Content = *payload.Content
	}

	if payload.Title != nil {
		post.Title = *payload.Title
	}

	if err := app.store.Posts.Update(request.Context(), post); err != nil {
		app.internalServerResponse(writer, request, err)
		return
	}

	if err := app.jsonResponse(writer, http.StatusOK, post); err != nil {
		app.internalServerResponse(writer, request, err)
		return
	}
}

func (app *application) postContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		idParam := chi.URLParam(request, "postId")
		postId, err := strconv.ParseInt(idParam, 10, 64)

		post, err := app.store.Posts.GetById(ctx, postId)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrorNotFound):
				app.notFoundResponse(writer, request, err)
			default:
				app.internalServerResponse(writer, request, err)
			}
			return
		}

		ctx = context.WithValue(ctx, postCtx, post)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func getPostFromCtx(r *http.Request) *store.Post {
	post, _ := r.Context().Value(postCtx).(*store.Post)
	return post
}
