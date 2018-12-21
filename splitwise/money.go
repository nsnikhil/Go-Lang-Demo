package main

type money int

func (m money) add(secondValue money) money {
	return m + secondValue
}

func (m money) subtract(secondValue money) money {
	return m - secondValue
}

func (m money) split(partitions int) money {
	return m / money(partitions)
}


