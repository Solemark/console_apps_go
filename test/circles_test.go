package console_apps

import (
	"math"
	ca "solemarc/go/console_apps/src"
	"testing"
)

func TestCircleArea(t *testing.T) {
	for _, r := range []float64{0.0, 1.0, 2.1, -3.2} {
		exp := 0.0
		if r > 0 {
			exp = math.Pi * (r * r)
		} else {
			exp = 0
		}
		res, e := ca.GetArea(r)
		if e != nil && res != 0 {
			t.Errorf("error thrown\nexpected: 0\nrecieved: %g", res)
		}
		if exp != res {
			t.Errorf("Error! Expected: %g\nrecieved: %g", exp, res)
		}
	}
}

func TestCircleCirc(t *testing.T) {
	for _, r := range []float64{0.0, 1.0, 2.1, -3.2} {
		exp := 0.0
		if r > 0 {
			exp = 2 * math.Pi * r
		} else {
			exp = 0
		}
		res, e := ca.GetCirc(r)
		if e != nil && res != 0 {
			t.Errorf("error thrown\nexpected: 0\nrecieved: %g", res)
		}
		if exp != res {
			t.Errorf("Error! Expected: %g\nrecieved: %g", exp, res)
		}
	}
}
