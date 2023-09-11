package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// make the getgreeting functions
func (eb englishBot) getGreeting() string {
	// custom logic very specific to english bot only
	return "Hi There Englishbot!"
}

func (sb spanishBot) getGreeting() string {
	// custom logic very specific to spanish bot only
	return "Hola Spanishbot!"
}

// redo the print greeting function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
