package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)

	colors["white"] = "#ffffff"

	delete(colors, "white")

	// colors := map[string]string{
	// 	"red":   "#cc0404",
	// 	"green": "#02ff11",
	// 	"blue":  "#1121ff",
	// }

	fmt.Println(colors)
}
