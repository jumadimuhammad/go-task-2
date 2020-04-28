package main

import (
	"fmt"
)

func main() {
	var result []int
	x := 7
	y := 70
	for i := x; i < y; i++ {
		temp := 0
		for j := 1; j <= i; j++ {
			if i <= 2 {
				result = append(result, i)

			} else if i%j == 0 {
				temp++
			}

		}
		if temp == 2 {
			result = append(result, i)
		}

	}
	fmt.Println(result)
}
