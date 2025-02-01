package console_apps

import (
	"sort"
)

func NumericalSort(data []int) []int {
	sort.Ints(data)
	return data
}
