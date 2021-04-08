package main

import (
	"fmt"
)

type Employee struct {
	name        string `required`
	phoneNumber string
	email       string
	age         int
}

type Vehicle struct {
	speed int
}

type Car struct {
	Vehicle
	color string
	make  int
}

func main() {
	populations := map[string]int{
		"Delhi":  100,
		"Mumbai": 120,
	}

	populations2 := make(map[string]int)
	populations2["Indore"] = 10

	fmt.Println(populations)
	fmt.Println(populations2)
	fmt.Println(populations["Delhi"])
	delete(populations, "Delhi")
	fmt.Println(populations["Delhi"], populations["Delhi2"])

	_, ok := populations["Delhi"]
	fmt.Printf("Value found %v\n\n", ok)

	employee := Employee{
		name:        "Sanil",
		phoneNumber: "123",
		age:         22,
		email:       "sanilkhurana7@gmail.com",
	}
	fmt.Println(employee, employee.name)

	employees := make([]Employee, 3)
	fmt.Println(employees)

	anonymousStruct := struct{ name string }{name: "Sanil"}
	fmt.Println(anonymousStruct)

	pointerToAnonymousStruct := &anonymousStruct
	fmt.Println(pointerToAnonymousStruct)

	anonymousStruct.name = "Not sanil"
	fmt.Println(anonymousStruct)
	fmt.Println(pointerToAnonymousStruct)

	fmt.Println("\n\n")

	c := Car{
		Vehicle: Vehicle{speed: 200},
		make:    1990,
		color:   "black",
	}
	c.speed = 100
	fmt.Println(c)

}
