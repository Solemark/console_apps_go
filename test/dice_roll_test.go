package console_apps

import (
	ca "solemarc/go/console_apps/src"
	"testing"
)

func TestDiceRoll(t *testing.T) {
	var dice []int = []int{4, 6, 8, 10, 12, 20}

	for _, roll := range dice {
		var answer int = ca.DiceRoll(roll)

		if (roll - answer) < 0 {
			t.Error("rolled a: d%i\nrecieved: %i", roll, answer)
		}
	}
}
