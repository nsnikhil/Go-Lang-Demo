package main

import (
	"math/rand"
	"time"
)

type number int

func (n number) add(secondNumber number) number {
	return n + secondNumber
}

func (n number) subtract(secondNumber number) number {
	return n - secondNumber
}

func (n number) divide(secondNumber number) number {
	return n / secondNumber
}

func (n number) multiply(secondNumber number) number {
	return n * secondNumber
}

func supplyNumber(limit int) number {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return number(r.Intn(limit))
}

func supplyNumberList(size int, limit int) []number {
	var arr []number
	for i := 0; i < size; i++ {
		arr = append(arr, supplyNumber(limit))
	}
	return arr
}

func getMax(firstNumber number, secondNumber number) number {
	if firstNumber > secondNumber {
		return firstNumber
	}
	return secondNumber
}

func getMin(firstNumber number, secondNumber number) number {
	if firstNumber < secondNumber {
		return firstNumber
	}
	return secondNumber
}
