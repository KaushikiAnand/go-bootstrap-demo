// Package handlers provides request handlers.
package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/KaushikiAnand/go-bootstrap-demo/project-templates/mysql/models"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func getCurrentUser(w http.ResponseWriter, r *http.Request) *models.UserRow {
	sessionStore := r.Context().Value("sessionStore").(sessions.Store)
	session, _ := sessionStore.Get(r, "$GO_BOOTSTRAP_PROJECT_NAME-session")
	return session.Values["user"].(*models.UserRow)
}

func getIdFromPath(w http.ResponseWriter, r *http.Request) (int64, error) {
	userIdString := mux.Vars(r)["id"]
	if userIdString == "" {
		return -1, errors.New("user id cannot be empty.")
	}

	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
