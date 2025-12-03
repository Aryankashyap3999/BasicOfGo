package subdir

import "fmt"

func Array () {
	var arr1 []int // slice :- dynamically type
	arr1 = append(arr1, 1, 2, 3, 4, 5)
	fmt.Println("Array 1: ", arr1)

	var arr2 [3]int // array
	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3
	fmt.Println("Array 2: ", arr2)

}