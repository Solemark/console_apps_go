package console_apps

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	FGO = "FGO"
	AK  = "AK"
	GI  = "GI"
)

func GachaRoll(game string) (string, error) {
	switch game {
	case FGO:
		return roll(100, 300, 5, "FGO"), nil
	case AK:
		return roll(50, 100, 6, "AK"), nil
	case GI:
		return roll(60, 90, 5, "GI"), nil
	default:
		return "", errors.New("unknown game")
	}
}

func roll(rate int, pity int, rarity int, game string) string {
	for roll := range pity + 1 {
		if rate == rand.Intn(pity+1) {
			return fmt.Sprintf("It took %d rolls to get a %d in %s", roll, rarity, game)
		}
	}
	return fmt.Sprintf("You hit pity at %d rolls for your %d in %s", pity, rarity, game)
}
