// Package weather provides a tool to forecast a location's current condition (weather).
package weather


var (
    // CurrentCondition represents the current condition of a city.
	CurrentCondition string
    // CurrentLocation represents the location of a city.
	CurrentLocation  string
)
// Forecast takes the location and then the condition and returns the current condition (weather) of said location (city).
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
