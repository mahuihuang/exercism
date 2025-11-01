// Package weather provides forecast the current weather condition of various cities in Goblinocus.
package weather

var (
	// CurrentCondition represents current weather.
	CurrentCondition string
	// CurrentLocation represents current city.
	CurrentLocation string
)

// Forecast return current condition of city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
