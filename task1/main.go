package main

import (
	"fmt"
)

func main() {
	array1 := [20]int{2, 3, 5, 1, 7, 23, 56, 21, 5, 6, 8, 9, 3, 2, 5, 12, 14, 65, 24, 13}
	array2 := [20]int{2, 31, 5, 1, 8, 43, 44, 21, 5, 6, 8, 9, 3, 7, 14, 12, 16, 65, 24, 13}

	var a1 []int
	var a2 []int
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			flag := false
			if array1[i] == array2[j] {
				flag = true
				// a1 = append(a1, array1[i])
				// a2 = append(a2, array2[j])
			}
			if !flag {
				a1 = append(a1, array1[i])
				a2 = append(a2, array2[j])
			}
		}
	}

	fmt.Println("a1 =", a1)
	fmt.Println("a2 =", a2)
}
