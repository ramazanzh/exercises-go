package problem11

func removeDups[T comparable](input []T) []T {
	var res []T
	temp := map[T]bool{}

	for _, value := range input {
		if temp[value] {
			continue
		}

		temp[value] = true
		res = append(res, value)
	}

	return res
}
