package mathproblems

import "math"

func canMakeArithmeticProgression(arr []int) bool {
	var seq int = 0

	for i := 0; i < len(arr); i++ {

		if i > 0 && seq != 0 {
			seq = arr[i] - arr[i-1]
			positive := math.Abs()
			if positive {

			}
		}
	}

}
