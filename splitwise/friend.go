package main

type friend struct {
	name string
	money
}

func (f *friend) addMoney(newMoney money) {
	(*f).money = (*f).add(newMoney)
}

func (f *friend) reduceMoney(newMoney money) {
	(*f).money = (*f).subtract(newMoney)
}


