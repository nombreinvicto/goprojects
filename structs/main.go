package main

import "fmt"

// contact info struct
type contactInfo struct {
	email   string
	zipCode int
}

// struct defn goes here
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// declare a value of type person
	alex := person{
		firstName: "Mahmud",
		lastName:  "Hasan",
		contact: contactInfo{
			"mahmud@hasan.com",
			83716,
		},
	}

	// go assigns zero values
	// to uninitialised values
	var ivy person
	fmt.Println(alex, ivy)

}
