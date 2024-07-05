package problem9

func factory(mult int) func(param ...int) []int {
	return func(param ...int) []int {
		var res []int
		for _, value := range param {
			res = append(res, mult*value)
		}

		return res
	}
}
