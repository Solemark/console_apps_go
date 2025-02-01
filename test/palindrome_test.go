package console_apps

import (
	ca "solemarc/go/console_apps/src"
	"testing"
)

func TestPalindrome(t *testing.T) {
	input := []string{"DAD", "Dad"}
	flag := []bool{true, false}

	for key, str := range input {
		var result bool = ca.Palindrome(str)
		if result != flag[key] {
			t.Errorf("\nexpected: %t\nrecieved: %t\nfor: %s = %t", flag[key], result, str, flag[key])
		}
	}
}
