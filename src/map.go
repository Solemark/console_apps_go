package console_apps

func Map[N Number](input []N, mapFunc func(N) N) []N {
	var output []N
	for _, item := range input {
		output = append(output, mapFunc(item))
	}
	return output
}
