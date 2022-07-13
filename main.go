package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Food struct {
	Name        string
	Cost, Stock int
}

type FoodMap map[string][]Food

var Foods = FoodMap{
	"snack": {
		{
			Name:  "tahu crispy",
			Cost:  15,
			Stock: 50,
		},
		{
			Name:  "roast potato",
			Cost:  17,
			Stock: 30,
		},
		{
			Name:  "potato wedges",
			Cost:  17,
			Stock: 32,
		},
		{
			Name:  "fluffy pancake",
			Cost:  20,
			Stock: 18,
		},
		{
			Name:  "french toast",
			Cost:  28,
			Stock: 12,
		},
	},
	"main course": {
		{
			Name:  "spaghetti aglio e olio",
			Cost:  28,
			Stock: 10,
		},
		{
			Name:  "spaghetti carbonara",
			Cost:  33,
			Stock: 8,
		},
		{
			Name:  "nasi goreng seafood",
			Cost:  24,
			Stock: 20,
		},
		{
			Name:  "mie goreng",
			Cost:  18,
			Stock: 28,
		},
		{
			Name:  "dori sambal matah",
			Cost:  23,
			Stock: 18,
		},
	},
	"coffee": {
		{
			Name:  "cold brew",
			Cost:  16,
			Stock: 21,
		},
		{
			Name:  "americano",
			Cost:  16,
			Stock: 13,
		},
		{
			Name:  "caramel macchiato",
			Cost:  18,
			Stock: 32,
		},
		{
			Name:  "espresso",
			Cost:  8,
			Stock: 12,
		},
		{
			Name:  "cappucino",
			Cost:  12,
			Stock: 16,
		},
		{
			Name:  "latte",
			Cost:  13,
			Stock: 11,
		},
	},
	"non coffee": {
		{
			Name:  "vanilla cookies",
			Cost:  17,
			Stock: 23,
		},
		{
			Name:  "matcha",
			Cost:  15,
			Stock: 8,
		},
		{
			Name:  "chocolate",
			Cost:  18,
			Stock: 12,
		},
		{
			Name:  "lemon tea",
			Cost:  8,
			Stock: 22,
		},
		{
			Name:  "thai tea",
			Cost:  11,
			Stock: 10,
		},
		{
			Name:  "air mineral",
			Cost:  5,
			Stock: 999,
		},
	},
}

func (food FoodMap) Add(category string, data Food) {
	Foods[category] = append(Foods[category], data)
}

func showMenu() {
	clearConsole()

	LINE_LENGTH := 50
	maxWord := getMaxWord(Foods)
	for key, value := range Foods {
		keyLength := (len(key) + 2) / 2
		fmt.Println(createLine("=", LINE_LENGTH/2-keyLength), strings.ToUpper(key), createLine("=", LINE_LENGTH/2-keyLength)+"\n")
		for _, food := range value {
			stock := strconv.Itoa(food.Stock)
			cost := strconv.Itoa(food.Cost)
			if len(cost) == 1 {
				cost = " " + cost
			}

			fmt.Println(strings.Title(food.Name) + createLine(" ", maxWord-len(food.Name)) + "\t" + cost + " K\t Stock : " + stock)
		}
		fmt.Println()
	}

	showOption()
}

func showOption() {
	fmt.Println()
	fmt.Println("1. Menambahkan menu baru")
	fmt.Println("2. Hapus menu")
	fmt.Println("3. Masukan pesanan")

	var ans string
	fmt.Printf("\nMasukan pilihan anda : ")
	fmt.Scanf("%s", &ans)

	clearConsole()
	handleUserInput(ans)
}

func handleUserInput(input string) {
	switch input {
	case "1":
		addMenu()
	// case "2":
	// 	removeMenu()
	// case "3":
	// 	addOrder()
	default:
		var ans string
		fmt.Printf("Invalid Input\n\nPress Enter to Continue...")
		fmt.Scanln(&ans)
		showMenu()
	}
}

func addMenu() {
	var category string
	var foodData Food

	fmt.Printf("Masukkan menu kategori : ")
	fmt.Scanf("%s", &category)
	fmt.Printf("Masukkan nama menu : ")
	fmt.Scanf("%s", &foodData.Name)
	fmt.Printf("Masukkan harga menu : ")
	fmt.Scanf("%d", &foodData.Cost)
	fmt.Printf("Masukkan stock yang tersedia : ")
	fmt.Scanf("%d", &foodData.Stock)

	Foods.Add(category, foodData)
	showMenu()

}

func main() {
	showMenu()
}

func createLine(lineType string, count int) string {
	line := ""
	for i := 0; i < count; i++ {
		line += lineType
	}
	return line
}

func getMaxWord(foodMap map[string][]Food) int {
	SPACE := 3
	result := 0
	for _, food := range foodMap {
		for _, f := range food {
			if result < len(f.Name) {
				result = len(f.Name) + SPACE
			}
		}
	}
	return result
}

func clearConsole() {
	fmt.Printf("\x1bc")
}
