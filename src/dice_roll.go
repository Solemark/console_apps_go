package console_apps

import (
	"math/rand"
)

func DiceRoll(max int) int {
	return rand.Intn(max) + 1
}
