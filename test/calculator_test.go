package console_apps_test

import (
	ca "solemarc/go/console_apps/src"
	"testing"
)

func getNumbers() [][]float32 {
	return [][]float32{
		{5.5, 5.5},
		{5.5, -5.5},
		{-5.5, -5.5},
	}
}

func TestAddition(t *testing.T) {
	var numbers [][]float32 = getNumbers()
	var expect float32
	var result float32

	for _, item := range numbers {
		expect = item[0] + item[1]
		result = ca.Add(item[0], item[1])

		if expect != result {
			t.Fatalf("expected: %f\n recieved: %f", expect, result)
		}
	}
}

func TestSubtraction(t *testing.T) {
	var numbers [][]float32 = getNumbers()
	var expect float32
	var result float32

	for _, item := range numbers {
		expect = item[0] - item[1]
		result = ca.Sub(item[0], item[1])

		if expect != result {
			t.Fatalf("expected: %f\n recieved: %f", expect, result)
		}
	}
}

func TestMultiplication(t *testing.T) {
	var numbers [][]float32 = getNumbers()
	var expect float32
	var result float32

	for _, item := range numbers {
		expect = item[0] * item[1]
		result = ca.Mul(item[0], item[1])

		if expect != result {
			t.Fatalf("expected: %f\n recieved: %f", expect, result)
		}
	}
}

func TestDivision(t *testing.T) {
	var numbers [][]float32 = getNumbers()
	var expect float32
	var result float32

	for _, item := range numbers {
		expect = item[0] / item[1]
		result = ca.Div(item[0], item[1])

		if expect != result {
			t.Fatalf("expected: %f\n recieved: %f", expect, result)
		}
	}
}
