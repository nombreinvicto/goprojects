package main

import (
	"fmt"
	"net/http"
	"os"
)

// declare global vars
var portNumber int = 3000

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
	_, err := fmt.Fprint(w, "<h1>Contact Page</h1><p>email:<a href=\"mailto:mhasan3.com\"")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// attach handler
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	// start the server
	serverString := fmt.Sprintf("http://localhost:%d/", portNumber)
	fmt.Println(serverString)
	err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber),
		nil)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
}
