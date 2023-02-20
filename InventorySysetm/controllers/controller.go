package controller

import (
	"fmt"
	mod "tsk/models"
)

var products = []mod.Product{
	{ID: 1, Name: "ABC", Price: 9.99, Quantity: 10},
	{ID: 2, Name: "BCD", Price: 19.99, Quantity: 5},
	{ID: 3, Name: "CEF", Price: 29.99, Quantity: 3},
}
var cart = mod.Cart{
	Prod: []mod.Product{
		{ID: 1, Name: "ABC", Price: 9.99, Quantity: 2},
		{ID: 2, Name: "BCD", Price: 19.99, Quantity: 5},
	},
	Total: 200,
}

func UpdateProd(name string, price float64, quantity int, id int) {
	for _, prod := range products {
		if prod.Name == name {
			prod.Quantity = quantity
			prod.Price = price
			prod.Name = name
			break
		}
	}
}
func UpdateQuantity(name string, quantity int) {
	for i, prod := range products {
		if prod.Name == name {
			products[i].Quantity = quantity
			break
		}
	}
}

func RemoveProd(name string) {

	var found bool
	fmt.Println()
	for i, prod := range products {
		if prod.Name == name {
			products = append(products[:i], products[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Product not found")
		return
	}

}
func AddNewProduct(name string, price float64, quantity int) {
	id := len(products) + 1
	newProd := mod.Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
	products = append(products, newProd)
}

func ReadAllProducts() {
	fmt.Println("-----------------")
	for _, prod := range products {
		fmt.Printf("| %s %.2f %d |\n", prod.Name, prod.Price, prod.Quantity)
	}
	fmt.Println("-----------------")

}
func ViewCart() {
	fmt.Println("----------Cart---------")
	for _, prod := range cart.Prod {
		fmt.Printf("| %s %.2f %d |\n", prod.Name, prod.Price, prod.Quantity)
	}
	fmt.Printf("| Total: %.2f |\n", cart.Total)
	fmt.Println("------------------------")
}
func AddToCart(name string, quantitySold int) {
	var total float64
	var found bool

	for i, prod := range products {
		if prod.Name == name && prod.Quantity >= quantitySold {
			products[i].Quantity -= quantitySold
			cart.Prod = append(cart.Prod, mod.Product{
				ID:       prod.ID,
				Name:     prod.Name,
				Price:    prod.Price,
				Quantity: quantitySold,
			})
			total += float64(quantitySold) * prod.Price
			found = true
			break
		}
	}

	if !found {
		fmt.Println("No product found with name: ", name)
		return
	}

	cart.Total += total
	fmt.Println("Added to cart sucessfully..")
}

func RemoveFromCart(name string) {

	var quant int
	var found bool
	for i, prod := range cart.Prod {
		if prod.Name == name {
			cart.Prod = append(cart.Prod[:i], cart.Prod[i+1:]...)
			quant = prod.Quantity
			found = true
			break
		}
	}
	if !found {
		fmt.Println("No product found in cart with name: ", name)
		return
	}

	//updating the qunatity
	for i, prod := range products {
		if prod.Name == name {
			products[i].Quantity += quant
			break
		}
	}

	//updating the total
	var total float64 = 0
	for _, prod := range cart.Prod {
		total += prod.Price * float64(prod.Quantity)
	}
	cart.Total = total

}
