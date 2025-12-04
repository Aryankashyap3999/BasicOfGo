package subdir

import "fmt"

// interface defination
type SellableProduct interface {
	buy() 
	getDiscount() int
}

type Product struct {
	name string
	price int
	company string
}

func newProduct(name string, price int, company string) *Product { // It resambles to a constructor
	p := Product{
		name: name,
		price: price,
		company: company,
	}
	return &p
}

// member function of struct
func (p *Product) displayInfo() {
	fmt.Println("Product Name:", p.name)
	fmt.Println("Product Price:", p.price)
	fmt.Println("Product Company:", p.company)
}

func (p *Product) buy() {
	fmt.Println("Buying product:", p.name, "from", p.company, "at price", p.price)
}

func (p *Product) getDiscount() int {
	// For demonstration, returning a fixed discount value
	return 35
}

func check_discount_and_buy (p SellableProduct) {
	discount := p.getDiscount()
	if discount > 30 {
		fmt.Println("Discount is good, buy the product");
		p.buy()
		return
	} else {
		fmt.Println("Discount is low, don't buy the product");
		return
	}
}
	

func Structs() {
	// product := Product{
	// 	name: "Laptop",
	// 	price: 1000,
	// 	company: "TechCorp",
	// }

	// fmt.Println("Product Name:", product.name)
	// fmt.Println("Product Price:", product.price)
	// fmt.Println("Product Company:", product.company)	

	new_p := newProduct("Laptop", 1000, "TechCorp")
	new_p.displayInfo()

	check_discount_and_buy(new_p)

}