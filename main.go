package main

import (
	"fmt"
	"log"

	"github.com/evad1n/vocab-scraper/interactive"
)

func main() {
	for !getAction() {
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
		if err := interactive.FindDefinitions(); err != nil {
			log.Fatalf("combining files: %v", err)
		}
	case 2:
		if err := interactive.CombineFiles(); err != nil {
			log.Fatalf("combining files: %v", err)
		}
	default:
		fmt.Printf("%d is not a valid choice\n\n", choice)
		return false
	}

	return true
}
