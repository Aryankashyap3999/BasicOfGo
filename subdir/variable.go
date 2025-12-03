package subdir

import "fmt"

func Variable() {
    fmt.Println("Hello from variable")

	const availableStock = 50
	var productName string
	var productPrize int = 10000
	var CompanyName = "Apple"
	productName = "Iphone"

	Category := "Electronice"

	fmt.Println("Stocks are: ", availableStock)
	fmt.Println("Product is ", productName, " and prise is ", productPrize, " category is", Category, " and company name is ", CompanyName);
}
