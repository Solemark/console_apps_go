package console_apps_test

import (
	ca "solemarc/go/console_apps/src"
	"strings"
	"testing"
)

func TestGachaRoll(t *testing.T) {
	var data []string = []string{
		"FGO", "AK", "GI",
	}

	for _, value := range data {
		var expect string = ca.GachaRoll(value)
		if !strings.Contains(expect, value) {
			t.Errorf("Expected: %s\nRecieved: %s", expect, value)
		}
	}
}
