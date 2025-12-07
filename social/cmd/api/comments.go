package main

import (
	"net/http"

	"github.com/vanhieuhp/social/internal/store"
)

type CreateCommentPayload struct {
	Content string `json:"content" validate:"required,max=10000"`
}

func (app *application) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateCommentPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err := Validate.Struct(payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := getPostFromCtx(r)

	comment := &store.Comment{
		Content: payload.Content,
		UserId:  post.UserID,
		PostId:  post.ID,
	}

	ctx := r.Context()

	if err := app.store.Comments.CreateComment(ctx, comment); err != nil {
		app.internalServerResponse(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerResponse(w, r, err)
		return
	}
}
