package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, d dir) [][]int {
	res := make([][]int, n)

	for index, _ := range res {
		res[index] = make([]int, n)

		for i := 0; i < n; i++ {
			res[index][i] = valueByIndex(d, n, index, i)
		}
	}

	return res
}

func valueByIndex(d dir, n int, si int, i int) int {
	if d == ul {
		return si + i
	}

	if d == ur {
		return n + si - 1 - i
	}

	if d == ll {
		return n - 1 - si + i
	}

	if d == lr {
		return (n * 2) - 2 - si - i
	}

	return 0
}
