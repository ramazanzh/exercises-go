package problem10

func factory() (map[string]int, func(name string) func(num int)) {
	brands := map[string]int{}

	return brands, func(name string) func(num int) {
		brands[name] = 0

		return func(num int) {
			brands[name] += num
		}
	}
}
