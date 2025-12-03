package subdir

import "fmt"

func Loops() {
	fmt.Println("Entered into the Loops")

	for i := 0; i < 5; i++ {
		fmt.Println("Loop iterations ", i)
	}

	for i := range 3 {
		fmt.Println("Range loop iterations ", i)
	}

	for i, char := range "Aryan" {
		fmt.Println("Range loop iterations ", i, char)
	}

	j := 10

	for j > 0 {
		if j == 3 {
			j--
			fmt.Println("Skipping iteration")
			continue
		}
		fmt.Println("While loop iterations ", j)
		j--;
	}

	for {
		fmt.Println("Infinite loop iterations ")
		break
	}
}