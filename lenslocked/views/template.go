// Package views is the V in our MVC pattern
// contains logic of everything related to rendering html
package views

import (
	"fmt"
	"html/template"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

// Execute method writes the template to response
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, nil)
	if err != nil {
		fmt.Printf("error executing template %v", err)
		http.Error(w, "error executing template",
			http.StatusInternalServerError)
		return
	}
}

// Parse parses template from file location
func Parse(filepath string) (Template, error) {
	// returns a *html.Tempalte
	htmlTpl, err := template.ParseFiles(filepath)
	if err != nil {
		// using fmt.Errorf to wrap err with additional context
		return Template{}, fmt.Errorf("error parsing template %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}
