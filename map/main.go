package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#4bf298",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["green"] = "4bf298"
	colors["red"] = "#ff0000"

	printMap(colors)

	delete(colors, "white")

	fmt.Println("After deleting white color")
	printMap(colors)
}

func printMap(c map[string]string){
	for color, hexCode := range c {
		fmt.Println(color +" , "+ hexCode)
	}
}