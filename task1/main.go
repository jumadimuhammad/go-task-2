package main

import (
	"fmt"
)

func main() {
	array1 := [20]int{2, 3, 5, 1, 7, 23, 56, 21, 5, 6, 8, 9, 3, 2, 5, 12, 14, 65, 24, 13}
	array2 := [20]int{2, 31, 5, 1, 8, 43, 44, 21, 5, 6, 8, 9, 3, 7, 14, 12, 16, 65, 24, 13}

	arr1 := []int{}
	arr2 := []int{}

	for i := 0; i < len(array1); i++ {
		flag := false
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				flag = true
			}
		}
		if !flag {
			arr1 = append(arr1, array1[i])

		}
	}

	for i := 0; i < len(array2); i++ {
		flag := false
		for j := 0; j < len(array1); j++ {
			if array2[i] == array1[j] {
				flag = true
			}
		}
		if !flag {
			arr2 = append(arr2, array2[i])

		}
	}

	fmt.Println("array1 =", arr1)
	fmt.Println("array2 =", arr2)
}
