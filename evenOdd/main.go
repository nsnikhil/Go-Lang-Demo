package main

import "fmt"

func main() {
	var arr []int

	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}

	for i, num := range arr {
		if num%2 == 0 {
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
	}
}
