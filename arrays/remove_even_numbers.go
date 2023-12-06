package arrays

func removeEvenNumbersFromArray(arr []int) (oddNumbers []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			oddNumbers = append(oddNumbers, arr[i])
		}
	}

	return
}
