package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/theycallmereza/social-app/internal/store"
	"net/http"
	"strconv"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
	ctx := r.Context()
	user, err := app.store.Users.GetByID(ctx, userID)

	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}
	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}
