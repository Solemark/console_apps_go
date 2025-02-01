package console_apps

import (
	ca "solemarc/go/console_apps/src"
	"strings"
	"testing"
)

func TestGachaRoll(t *testing.T) {
	data := []string{
		"FGO", "AK", "GI", "unknown game",
	}

	for _, inp := range data {
		res, e := ca.GachaRoll(inp)
		if e != nil && res != "" {
			t.Errorf("recieved error\nexpected: \"\"\nrecieved: %s", res)
		} else if !strings.Contains(res, inp) && res != "" {
			t.Errorf("Expected: %s\nRecieved: %s", inp, res)
		}
	}
}
