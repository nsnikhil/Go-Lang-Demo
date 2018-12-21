package main_test

import (
	"awesomeProject/testing"
	"fmt"
)

func ExampleInts() {
	s := main.Ints(1, 2, 3, 4, 5)
	fmt.Println("sum of one to five is", s)
	//Output:
	//sum of one to five is 15
}
