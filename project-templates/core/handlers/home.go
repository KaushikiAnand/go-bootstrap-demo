package handlers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/KaushikiAnand/go-bootstrap-demo/project-templates/core/libhttp"
)

func main() {
	os.Getenv("GO_BOOTSTRAP_REPO_NAME")
	os.Getenv("GO_BOOTSTRAP_REPO_USER")
	os.Getenv("GO_BOOTSTRAP_PROJECT_NAME")
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, nil)
}
