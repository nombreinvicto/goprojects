package main

import (
	"fmt"
	"net/http"
	"os"
)

// declare global vars
var portNumber int = 3000

// listen and serve func requires
// router that implement ServeHttp method
type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// creating a basic router by ourself
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

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

	// start the server
	var router Router
	serverString := fmt.Sprintf("http://localhost:%d/", portNumber)
	fmt.Println(serverString)

	// pass nil to use default serve mux, mux = router
	err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), router)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
}
