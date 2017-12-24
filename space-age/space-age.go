package space

type Planet string

const earthSeconds = 31557600

func Age(seconds float64, planet Planet) float64 {

	switch planet {
	case "Mercury":
		return seconds / YearsToSeconds(0.2408467)
	case "Venus":
		return seconds / YearsToSeconds(0.61519726)
	case "Mars":
		return seconds / YearsToSeconds(1.8808158)
	case "Jupiter":
		return seconds / YearsToSeconds(11.862615)
	case "Saturn":
		return seconds / YearsToSeconds(29.447498)
	case "Uranus":
		return seconds / YearsToSeconds(84.016846)
	case "Neptune":
		return seconds / YearsToSeconds(164.79132)
	default:
		return seconds / YearsToSeconds(1)
	}
}

func YearsToSeconds(year float64) float64 {
	return year * earthSeconds
}
