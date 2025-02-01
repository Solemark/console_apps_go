package console_apps_test

import (
	"reflect"
	ca "solemarc/go/console_apps/src"
	"testing"
)

func TestEvens(t *testing.T) {
	expect := []int32{-4, -2, 0, 2, 4}
	data := []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	result := ca.Even(data)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}

func TestOdds(t *testing.T) {
	expect := []int32{-5, -3, -1, 1, 3, 5}
	data := []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	result := ca.Odd(data)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}

func TestNegative(t *testing.T) {
	expect := []int32{-5, -4, -3, -2, -1}
	data := []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	result := ca.Negative(data)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}

func TestPositive(t *testing.T) {
	expect := []int32{1, 2, 3, 4, 5}
	data := []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	result := ca.Positive(data)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}
