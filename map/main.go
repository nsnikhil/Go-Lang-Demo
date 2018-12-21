package main

import "fmt"

func main() {
	//color := make(map[string]string)
	color := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#000000",
	}
	color["grey"] = "#eoeoeo"
	printMap(color)
}

func printMap(m map[string]string) {
	for i, e := range m {
		fmt.Println("hex code for ", i, " is ", e)
	}
}
