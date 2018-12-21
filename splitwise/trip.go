package main

import "fmt"

type trip []friend

func generateTrip() trip {
	return trip([]friend{{name: "first", money: 120},
		{name: "second", money: 120},
		{name: "third", money: 360}})
}

func (t trip) getTotalMoney() money {
	var total money = 0
	for _, e := range t {
		total = e.money.add(total)
	}
	return total
}

func (t trip) getAverageExpenditure() money {
	return t.getTotalMoney().split(len(t))
}

func (t trip) printResult() {
	for _, e := range t {
	inner:
		for _, f := range t {
			if e.name != f.name {
				if e.money > t.getAverageExpenditure() && f.money < t.getAverageExpenditure() {
					fmt.Println(f.name,
						" gives ",
						t.getAverageExpenditure().subtract(f.money),
						" to ",
						e.name)
					continue inner
				}
			}
		}
	}
}

