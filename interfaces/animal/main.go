package main

func main() {
	var animals []animal

	animals = append(animals, dog{})
	animals = append(animals, cat{})

	for _, animal := range animals {
		printAnimal(animal)
	}


}

func printAnimal(a animal) {
	a.eat()
	a.sleep()
	a.walk()
}
