package handlers

func main() {
    repoName := os.Getenv("GO_BOOTSTRAP_REPO_NAME")
    repoUser := os.Getenv("GO_BOOTSTRAP_REPO_USER")
    projectName := os.Getenv("GO_BOOTSTRAP_PROJECT_NAME")
}

import (
	"$GO_BOOTSTRAP_REPO_NAME/$GO_BOOTSTRAP_REPO_USER/$GO_BOOTSTRAP_PROJECT_NAME/libhttp"
	"html/template"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, nil)
}
