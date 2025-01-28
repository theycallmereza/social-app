package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server Error: %s Path: %s Error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "An internal server error has occurred")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad Request Error: %s Path: %s Error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Conflict Error: %s Path: %s Error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not Found Error: %s Path: %s Error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusNotFound, "Not Found")
}
