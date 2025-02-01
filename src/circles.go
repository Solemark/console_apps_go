package console_apps

import (
	"errors"
	"math"
)

func GetArea[R Radius](radius R) (R, error) {
	if radius > 0 {
		return (math.Pi * (radius * radius)), nil
	} else {
		return 0, errors.New("radius <= 0")
	}
}

func GetCirc[R Radius](radius R) (R, error) {
	if radius > 0 {
		return (2 * math.Pi * radius), nil
	} else {
		return 0, errors.New("radius <= 0")
	}
}
