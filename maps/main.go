package main

import "fmt"

func main() {
	//[string - keys] string - values
	// var colors map[string]string
	// colors := map[string]string {
	// 	"red": "#FF0000",
	// 	"green": "#00FF00",
	// }
	// colors := make(map[string]string) // same, empty

	// colors["blue"] = "#0000FF"
	// fmt.Println(colors)
	// delete(colors, "blue")

	// fmt.Println(colors)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
	}
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
