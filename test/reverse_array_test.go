package console_apps

import (
	"reflect"
	ca "solemarc/go/console_apps/src"
	"testing"
)

func TestReverseArray(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := ca.ReverseArray(input)

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Expect: %d\nResult: %d", result, expect)
		return
	}
}
