package console_apps

import (
	"reflect"
	ca "solemarc/go/console_apps/src"
	"testing"
)

func TestMap(t *testing.T) {
	expect := []int32{3, 6, 9}
	result := ca.Map([]int32{1, 2, 3}, timesThree)
	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}

func timesThree[N ca.Number](x N) N {
	return x * 3
}
