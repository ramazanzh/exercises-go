package problem1

func isChangeEnough(pocket [4]int, price float32) bool {
	var total float32
	coinValues := [4]float32{0.25, 0.10, 0.05, 0.01}

	for index, coinValue := range coinValues {
		total += coinValue * float32(pocket[index])

		if total >= price {
			return true
		}
	}

	return false
}
