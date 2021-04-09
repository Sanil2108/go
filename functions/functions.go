package main

import "fmt"

func sayMessage(msg string) (string, error) {
	fmt.Println(msg)

	return "I said it", nil
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

type Animal struct {
	class string
	name  string
}

func (a Animal) printAnimal() {
	fmt.Printf("Name of the animal - %v, Class of the animal - %v", a.name, a.class)
}

func main() {
	didItSayIt, err := sayMessage("Hello world")

	fmt.Println(didItSayIt, err)

	a, b := 1, 2

	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

	anon := func(msg string) {
		fmt.Printf("\nThis is from an anonymous function %v", msg)
	}
	anon("ANONNNN")

	myAnimal := Animal{
		name:  "Koshka",
		class: "cat",
	}
	myAnimal.printAnimal()
}
