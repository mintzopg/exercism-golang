package space

const earth = 31557600

// Planet wrapper type
type Planet string

// Age function(uint)float64; calculate space age
func Age(secs float64, planet Planet) float64 {

	var age = -1.0
	switch planet {
	case "Earth":
		age = secs / earth
	case "Mercury":
		age = secs / (0.2408467 * earth)
	case "Venus":
		age = secs / (0.61519726 * earth)
	case "Mars":
		age = secs / (1.8808158 * earth)
	case "Jupiter":
		age = secs / (11.862615 * earth)
	case "Saturn":
		age = secs / (29.447498 * earth)
	case "Uranus":
		age = secs / (84.016846 * earth)
	case "Neptune":
		age = secs / (164.79132 * earth)
	}
	return age
}
