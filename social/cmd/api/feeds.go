package main

import "net/http"

func (app *application) getUserFeedHandler(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(2))
	if err != nil {
		app.internalServerResponse(writer, request, err)
		return
	}

	if err = app.jsonResponse(writer, http.StatusOK, feed); err != nil {
		app.internalServerResponse(writer, request, err)
		return
	}
}
