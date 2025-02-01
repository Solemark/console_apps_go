package console_apps

func Even(data []int32) []int32 {
	output := []int32{}
	for _, v := range data {
		if v%2 == 0 {
			output = append(output, v)
		}
	}
	return output
}

func Odd(data []int32) []int32 {
	output := []int32{}
	for _, v := range data {
		if v%2 != 0 {
			output = append(output, v)
		}
	}
	return output
}

func Negative(data []int32) []int32 {
	output := []int32{}
	for _, v := range data {
		if v < 0 {
			output = append(output, v)
		}
	}
	return output
}

func Positive(data []int32) []int32 {
	output := []int32{}
	for _, v := range data {
		if v > 0 {
			output = append(output, v)
		}
	}
	return output
}
