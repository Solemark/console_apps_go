package console_apps

import (
	ca "solemarc/go/console_apps/src"
	"testing"
)

func TestDatetime(t *testing.T) {
	expect := "the date is Friday the 26th of May 2023"
	answer := ca.CheckDate()

	if expect != answer {
		t.Errorf("expected: %s\n recieved: %s", expect, answer)
	}
}
