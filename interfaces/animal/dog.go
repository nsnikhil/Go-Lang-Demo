package main

import "fmt"

type dog struct{}

func (d dog) eat() {
	fmt.Println("dog eat")
}

func (d dog) sleep() {
	fmt.Println("dog sleep")
}

func (d dog) walk() {
	fmt.Println("dog walk")
}


