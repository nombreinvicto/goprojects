package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(visitorName string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s",
		greeter.LanguageName(), greeter.Greet(visitorName))
}

// implement the Italian struct and methods
type Italian struct{}

func (it Italian) LanguageName() string {
	return "Italian"
}

func (it Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

// implement the Prtuguese struct and methods
type Portuguese struct{}

func (pg Portuguese) LanguageName() string {
	return "Portuguese"
}

func (pg Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
