package spaceage

// Planet is a custom type to ensure type safety when looking up orbital periods.
type Planet string

const secondsInEarthYear = 31557600

// secondsPerPlanetYear maps each planet to its orbital period converted into seconds.
// This is defined at the package level to prevent re-allocation on every function call.
var secondsPerPlanetYear = map[Planet]float64{
	"Mercury": 0.2408467 * secondsInEarthYear,
	"Venus":   0.61519726 * secondsInEarthYear,
	"Earth":   1.0 * secondsInEarthYear,
	"Mars":    1.8808158 * secondsInEarthYear,
	"Jupiter": 11.862615 * secondsInEarthYear,
	"Saturn":  29.447498 * secondsInEarthYear,
	"Uranus":  84.016846 * secondsInEarthYear,
	"Neptune": 164.79132 * secondsInEarthYear,
}

// Age calculates how many years old someone would be on a given planet based on total seconds lived.
// It uses the pre-calculated secondsPerPlanetYear map to ensure O(1) lookup and minimal math at runtime.
func Age(seconds float64, planet Planet) float64 {
	// Pull the planet's year length in seconds. 
	// If the planet exists in the map, perform the division.
	planetSeconds, ok := secondsPerPlanetYear[planet]
	if !ok {
		return -1
	}

	return seconds / planetSeconds
}