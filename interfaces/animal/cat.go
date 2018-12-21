package main

import "fmt"

type cat struct{}

func (d cat) eat() {
	fmt.Println("cat eat")
}

func (d cat) sleep() {
	fmt.Println("cat sleep")
}

func (d cat) walk() {
	fmt.Println("cat walk")
}