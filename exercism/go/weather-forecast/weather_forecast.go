// Package weather provides functions to forecast weather.
package weather

// CurrentCondition is a string that describes the current condition of a place.
var CurrentCondition string

// CurrentLocation is a string that desribes current location of a place.
var CurrentLocation string

// Forecast takes city and condition as input strings
// and then returns a formatted string displaying the
// condition of the city in a nice manner.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
