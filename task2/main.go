package main

import (
	"fmt"
)

func main() {
	var genap []int
	var ganjil []int
	var prima []int

	array := [30]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 14, 15, 16, 12, 18, 19, 20, 22, 24, 26}

	for i := 0; i < len(array); i++ {
		if array[i] != 0 {

			if array[i]%2 == 0 {
				genap = append(genap, array[i])
			}
			if array[i]%2 != 0 {
				ganjil = append(ganjil, array[i])
			}

			if array[i] <=2{
				prima = append(prima, array[i])
			}else{
				temp := 0
				for j := 1; j <= array[i]; j++{
					if array[i] % j == 0 {
						temp ++
					}
				}
				if temp == 2{
					prima = append(prima, array[i])
				}
			}

		}
	}
	fmt.Println("Genap : ", genap, ", Count : ", len(genap))
	fmt.Println("Ganjil : ", ganjil, ", Count : ", len(ganjil))
	fmt.Println("Prima : ", prima, ", Count : ", len(prima))

}
