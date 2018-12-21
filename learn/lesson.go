package learn

import "fmt"

var deckSize int

func someFucntion() {
	deckSize = 50
	fmt.Println(deckSize)

	cards := []string{"card one", "card two", getCard()}

	fmt.Println(cards)
	fmt.Println(sliceTest())
	iterateSlice(sliceTest())
}

func variableDeclaration() {
	var something string = "someVariable"
	var something2 = "someVariable2"
	something3 := "someVariable3"
	fmt.Println(something, something2, something3)
}

func getCard() string {
	return "some card"
}

func returnSum() int {
	return 5
}

func sliceTest() []int {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	modifiedList := append(list, 0)
	fmt.Println(modifiedList)
	return list
}

func iterateSlice(list []int) {
	for i, item := range list {
		fmt.Println(i,item)
	}
}

