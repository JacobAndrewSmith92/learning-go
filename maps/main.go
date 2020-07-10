package main

import "fmt"

func main() {
	// 1. var colors map[string]string
	// 2. colors := map[string]string
	// 3. colors := make([string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#423fs0",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
