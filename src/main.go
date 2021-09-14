package main

import (
	"fmt"

	"github.com/evad1n/vocab-scraper/combine"
	"github.com/evad1n/vocab-scraper/define"
)

func main() {
	for getAction() {
	}
}

func getAction() bool {
	fmt.Println("What would you like to do:")
	fmt.Println("1: Define")
	fmt.Println("2: Combine")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		define.FindDefinitions()
	case 2:
		combine.CombineFiles()
	default:
		fmt.Printf("%d is not a valid choice\n", choice)
		return false
	}
	return true
}
