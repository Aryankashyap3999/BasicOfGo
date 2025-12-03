package subdir

import "fmt"

func Maps () {
	fmt.Println("Introduction to maps")

	productprices := map[string]int {
		"iphone 15 pro": 1000,
		"iphone 15": 800,
		"iphone 14": 600,
	}

	fmt.Println("Product prices: ", productprices)

	customMap := map[string]string{}

	fmt.Println("Custom Map ", customMap)

	emptyMap := make(map[string]int)

	fmt.Println("Empty Map: ", emptyMap)

	emptyMap["Key1"] = 100
	emptyMap["Key2"] = 200

	fmt.Println("Empty Map: ", emptyMap)

	for key, value := range emptyMap {
		fmt.Println("Key is ", key, " value is ", value)
	}

}