package subdir

import "fmt"

func Pointers () {
	i := 120

	var ptr *int = &i

	ptr1 := &i

	fmt.Println("Value of i: ", i, " Pointer to i: ", ptr, " Pointer1 to i: ", ptr1)
	fmt.Println("Value of pointer: ", *ptr1)
}