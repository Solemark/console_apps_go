package console_apps

import (
	"reflect"
	ca "solemarc/go/console_apps/src"
	"testing"
)

func TestNumericalSort(t *testing.T) {
	data := []int{7, 6, 5, 1, 8, 4, 9, 2, 3, 10}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := ca.NumericalSort(data)

	if !reflect.DeepEqual(exp, result) {
		t.Errorf("Expected: %d\nResult: %d", exp, result)
	}

}
