package templates

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := "templates/" + tmpl + ".html"
	t, err := template.ParseFiles(
		"templates/base.html",
		"templates/navbar.html",
		"templates/footer.html",
		tmplPath)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
	}
}
