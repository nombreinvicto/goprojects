// Package controllers contains all
// logic of controlling/handling requests
package controllers

import (
	"lenslocked/views"
	"net/http"
)

type Static struct {
	Template views.Template
}

// we are now making static implement http.Handler
func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}
