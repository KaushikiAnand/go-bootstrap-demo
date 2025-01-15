// Package handlers provides request handlers.
package handlers

func main() {
    repoName := os.Getenv("GO_BOOTSTRAP_REPO_NAME")
    repoUser := os.Getenv("GO_BOOTSTRAP_REPO_USER")
    projectName := os.Getenv("GO_BOOTSTRAP_PROJECT_NAME")
}

import (
	"errors"
	"$GO_BOOTSTRAP_REPO_NAME/$GO_BOOTSTRAP_REPO_USER/$GO_BOOTSTRAP_PROJECT_NAME/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

func getCurrentUser(w http.ResponseWriter, r *http.Request) *models.UserRow {
	sessionStore := r.Context().Value( "sessionStore").(sessions.Store)
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
