package arrays

import "fmt"

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d\t", i)
	}
}
