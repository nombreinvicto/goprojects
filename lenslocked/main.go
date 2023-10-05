package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// declare global vars
var portNumber int = 3000

// homeHandler is the handler func for homepage
// http.Request is what our browser sends to the server
// http.ResponseWriter used to write a response to the request
func homeHandler(w http.ResponseWriter, r *http.Request) {
	templateFilepath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, templateFilepath)
}

// handler for the contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	templateFilepath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, templateFilepath)
}

// common execute template func
func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		fmt.Printf("error parsing template %v", err)
		http.Error(w, "error parsing template",
			http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		fmt.Printf("error executing template %v", err)
		http.Error(w, "error executing template",
			http.StatusInternalServerError)
		return
	}
}

func main() {

	// instantiate chi router
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "Page Not Found", http.StatusNotFound)
	})

	// pass nil to use default serve mux, mux = router
	serverString := fmt.Sprintf("http://localhost:%d/", portNumber)
	fmt.Println(serverString)
	err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), r)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
}
