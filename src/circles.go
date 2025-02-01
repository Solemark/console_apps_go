package console_apps

import "math"

func GetArea[R Radius](radius R) R {
	switch {
	case radius > 0:
		return math.Pi * (radius * radius)
	default:
		return 0
	}
}

func GetCirc[R Radius](radius R) R {
	switch {
	case radius > 0:
		return 2 * math.Pi * radius
	default:
		return 0
	}
}
