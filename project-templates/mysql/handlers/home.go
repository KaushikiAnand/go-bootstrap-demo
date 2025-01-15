package handlers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/KaushikiAnand/go-bootstrap-demo/project-templates/core/libhttp"
	"github.com/KaushikiAnand/go-bootstrap-demo/project-templates/mysql/models"

	"github.com/gorilla/sessions"
)

func main() {
	os.Getenv("GO_BOOTSTRAP_REPO_NAME")
	os.Getenv("GO_BOOTSTRAP_REPO_USER")
	os.Getenv("GO_BOOTSTRAP_PROJECT_NAME")
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	sessionStore := r.Context().Value("sessionStore").(sessions.Store)

	session, _ := sessionStore.Get(r, "$GO_BOOTSTRAP_PROJECT_NAME-session")
	currentUser, ok := session.Values["user"].(*models.UserRow)
	if !ok {
		http.Redirect(w, r, "/logout", 302)
		return
	}

	data := struct {
		CurrentUser *models.UserRow
	}{
		currentUser,
	}

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, data)
}
