package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

// declare global vars
var portNumber int = 3000

// homeHandler is the handler func for homepage
// http.Request is what our browser sends to the server
// http.ResponseWriter used to write a response to the request
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// set the response header content type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprint(w, "<h1>Welcome to my awesome site!!!</h1>")
	if err != nil {
		fmt.Println(err)
	}
}

// handler for the contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprint(w, "<h1>Contact Page</h1><p>email:"+
		"<a href=\"mailto:mhasan3.com\"")
	if err != nil {
		fmt.Println(err)
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
