package console_apps

import (
	"reflect"
	ca "solemarc/go/console_apps/src"
	"testing"
)

func TestReverseArray(t *testing.T) {
	inp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	exp := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	res := ca.ReverseArray(inp)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("Expect: %d\nResult: %d", res, exp)
		return
	}
}
