package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting server on :3000")
	http.ListenAndServe(":3000", nil)
}

// http.Request is what our browser sends to the server
// http.ResponseWriter used to write a response to the request
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}
