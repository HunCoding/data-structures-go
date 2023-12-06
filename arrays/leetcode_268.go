package arrays

func findMissingNumber(arr []int) int {
	arrLength := len(arr) + 1
	sum := arrLength * (arrLength + 1) / 2
	for _, value := range arr {
		sum -= value
	}

	return sum
}
