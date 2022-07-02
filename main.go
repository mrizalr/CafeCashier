package main

import (
	"fmt"
	"strings"
)

var Foods = map[string][]string{
	"snack":       {"tahu crispy", "roast potato", "potato wedges", "fluffy pancakes", "french toast"},
	"main course": {"spaghetti aglio e olio", "spaghetti carbonara", "nasi goreng seafood", "mie goreng", "dori sambal matah"},
	"coffee":      {"cold brew", "americano", "caramel macchiato", "espresso", "cappucino", "latte"},
	"non coffee":  {"vanilla cookies", "matcha", "chocolate", "lemon tea", "thai tea", "air mineral"},
}

func showMenu() {
	for key, value := range Foods {
		fmt.Println("======", strings.Title(key), "======")
		for _, food := range value {
			fmt.Println(strings.Title(food))
		}
		fmt.Println()
	}
}

func main() {
	showMenu()
}
