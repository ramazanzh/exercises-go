package problem1

const (
	quarter float32 = 0.25
	dime    float32 = 0.1
	nickel  float32 = 0.05
	penny   float32 = 0.01
)

func isChangeEnough(pocket [4]int, price float32) bool {
	var total float32
	coins := [4]float32{quarter, dime, nickel, penny}

	for index, coinValue := range coins {
		total += coinValue * float32(pocket[index])

		if total >= price {
			return true
		}
	}

	return false
}
