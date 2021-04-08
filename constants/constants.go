package main

import (
	"fmt"
)

const (
	noAnimal = iota
	cat
	dog
	monkey
)

func main() {
	// A constant is declared with const
	const myConst string = "Hello world"

	// constants have to be known at compile time so cannot calculate them at runtime
	// const sinValue = math.Sin(1)

	// myConst = "HI"

	fmt.Println(myConst)
	// fmt.Println(sinValue)

	fmt.Println(noAnimal)
	fmt.Println(cat)
	fmt.Println(dog)
	fmt.Println(monkey)

	var animal1 int
	fmt.Println(animal1 == noAnimal)
	fmt.Println(animal1 == cat)
	fmt.Println(animal1 == dog)
	fmt.Println(animal1 == monkey)

	fmt.Println("\n\n")

	var animal2 int = cat
	fmt.Println(animal2 == noAnimal)
	fmt.Println(animal2 == cat)
	fmt.Println(animal2 == dog)
	fmt.Println(animal2 == monkey)
}
