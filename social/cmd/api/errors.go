package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerResponse(writer http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal server error: %s, path: %s, error: %s", r.Method, r.URL.Path, err)

	err = writeJSONError(writer, http.StatusInternalServerError, "the server encountered a problem")
	if err != nil {
		return
	}
}

func (app *application) notFoundResponse(writer http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not found error: %s, path: %s, error: %s", r.Method, r.URL.Path, err)

	err = writeJSONError(writer, http.StatusNotFound, "Resource not found")
	if err != nil {
		return
	}
}

func (app *application) badRequestResponse(writer http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad request error: %s, path: %s, error: %s", r.Method, r.URL.Path, err)
	err = writeJSONError(writer, http.StatusBadRequest, err.Error())
	if err != nil {
		return
	}
}

func (app *application) ConflictResponse(writer http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Conflict error: %s, path: %s, error: %s", r.Method, r.URL.Path, err)
	err = writeJSONError(writer, http.StatusConflict, err.Error())
	if err != nil {
		return
	}
}
