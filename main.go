package main

import "fmt"

var snack = [5]string{"Tahu cripsy", "Roast Potato", "Potato Wedges", "Fluffy Pancake", "French Toast"}
var mainCourse = [5]string{"Spaghetti Aglio E Olio", "Spaghetti Carbonara", "Butter Rice", "Nasi Goreng Seafood", "Mie Merah"}
var coffee = [8]string{"Cold Brew", "Americano", "Espresso", "Hazelnut", "Latte", "Tubruk", "Cappucino", "Kopi Gula Aren"}
var nonCoffee = [6]string{"Matcha", "Chocolate", "Lemon Tea", "Thai Tea", "Mint Tea", "Cheese Cookies"}

func main() {
	showMenu()
}

func showMenu() {
	fmt.Println("\n====== SNACK ======")
	for _, element := range snack {
		fmt.Println("-", element)
	}

	fmt.Println("\n====== MAIN COURSE ======")
	for _, element := range mainCourse {
		fmt.Println("-", element)
	}

	fmt.Println("\n====== COFFEE ======")
	for _, element := range coffee {
		fmt.Println("-", element)
	}

	fmt.Println("\n====== NON COFFEE ======")
	for _, element := range nonCoffee {
		fmt.Println("-", element)
	}
}
