package console_apps

import (
	"reflect"
	ca "solemarc/go/console_apps/src"
	"testing"
)

func TestMap(t *testing.T) {
	exp := []int32{3, 6, 9}
	res := ca.Map([]int32{1, 2, 3}, timesThree)
	if !reflect.DeepEqual(exp, res) {
		t.Error("Expected: ", exp, "\nRecieved: ", res)
	}
}

func timesThree[N ca.Number](x N) N {
	return x * 3
}
