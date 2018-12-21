package main

import (
	"fmt"
)

type number float64

//noinspection GoUnhandledErrorResult
func main() {
	var a int
	var b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Println(number(a).add(number(b)))
}

func (n number) add(o number) number {
	return n + o
}

func (n number) subtract(o number) number {
	return n - o
}

func (n number) multiply(o number) number {
	return n * o
}

func (n number) divide(o number) number {
	return n / o
}
