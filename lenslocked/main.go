package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"lenslocked/controllers"
	"lenslocked/views"
	"net/http"
	"os"
	"path/filepath"
)

// declare global vars
var portNumber int = 3000

func main() {

	// instantiate chi router
	r := chi.NewRouter()

	// goal: parse and check template integrity before even starting server
	// design philosophy = get rendered templates from views
	// then allow controllers to send the rendered templates
	// r.Method allows us to accept http.Hanlder
	// r.Get only accepts http.HandlerFunc
	homeTpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	r.Method(http.MethodGet, "/", controllers.Static{Template: homeTpl})
	if err != nil {
		panic(err)
	}

	contactTpl, err := views.Parse(filepath.Join("templates", "contact.gohtml"))
	r.Method(http.MethodGet, "/contact", controllers.Static{Template: contactTpl})
	if err != nil {
		panic(err)
	}

	// finally handling the nofound page
	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "Page Not Found", http.StatusNotFound)
	})

	// pass nil to use default serve mux, mux = router
	serverString := fmt.Sprintf("http://localhost:%d/", portNumber)
	fmt.Println(serverString)
	err = http.ListenAndServe(fmt.Sprintf(":%d", portNumber), r)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
}
