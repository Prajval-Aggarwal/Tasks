package main

import (
	"fmt"
	controller "tsk/controllers"
)

func main() {

	for {
		fmt.Println("----------------------Inventory Menu:---------------------")
		fmt.Println("                   1. View products")
		fmt.Println("                   2. Update product")
		fmt.Println("                   3. Remove product")
		fmt.Println("                   4. Add new product")
		fmt.Println("                   5. Remove a product")
		fmt.Println("                   6. View my cart")
		fmt.Println("                   7. Add products to cart")
		fmt.Println("                   8. Remove products from cart")
		fmt.Println("                   9. Quit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			fmt.Println("Viewing products...")
			controller.ReadAllProducts()

		case 2:
			fmt.Println("Update product...")
			var name string
			var price float64
			var quantity int
			var id int
			fmt.Print("Enter product name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter product price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter product quantity: ")
			fmt.Scanln(&quantity)
			fmt.Print("Enter product id: ")
			fmt.Scanln(&id)
			controller.UpdateProd(name, price, quantity, id)
			fmt.Println("Product added successfully!.....")
			controller.ReadAllProducts()

		case 3:
			fmt.Println("Removing product...")
			fmt.Println("Enter the name of the product you want to delete")
			var name string
			fmt.Scan(&name)
			controller.RemoveProd(name)
			fmt.Println("Product removed successfully!.....")
			controller.ReadAllProducts()

		case 4:
			fmt.Println("Adding new Product ...")
			var name string
			var price float64
			var quantity int
			fmt.Print("Enter product name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter product price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter product quantity: ")
			fmt.Scanln(&quantity)
			controller.AddNewProduct(name, price, quantity)
			fmt.Println("Product added successfully!.....")
			controller.ReadAllProducts()

		case 5:
			fmt.Println("Removing product")
			var name string
			fmt.Scanln(&name)
			controller.RemoveProd(name)

		case 6:
			fmt.Println("Viewing Yout Cart...")
			controller.ViewCart()

		case 7:
			fmt.Println("Adding products to your cart")

			fmt.Println("List of products that are in inventory are..")
			controller.ReadAllProducts()

			fmt.Println("Enter the product name you want to add to your cart:")
			var name string
			var quantity int
			fmt.Scanln(&name)
			fmt.Println("Enter the quantity you bought: ")
			fmt.Scanln(&quantity)

			controller.AddToCart(name, quantity)
			controller.ViewCart()

		case 8:
			fmt.Println("Removing product from cart....")
			var name string
			fmt.Println("Enter the name of the product you want to delete...")
			fmt.Scanln(&name)
			controller.RemoveFromCart(name)
			controller.ViewCart()

		case 9:
			fmt.Println("Quitting the program...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
