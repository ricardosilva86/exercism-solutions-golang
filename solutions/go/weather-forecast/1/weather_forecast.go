// Package weather does something related to the weather
// This is a comment to explain about this package.
package weather

// CurrentCondition variable describes the current weather condition.
var CurrentCondition string

// CurrentLocation variable describes the current location.
var CurrentLocation string

// Forecast function return a string that explains the current weather condition at the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
