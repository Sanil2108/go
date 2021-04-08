package main

import (
	"fmt"
)

func main() {
	grades := [3]int{1, 2, 3}

	grades = [...]int{1, 2, 3}

	var students [3]int

	fmt.Println(grades)

	fmt.Println(students)

	students = [...]int{10, 20, 30}

	fmt.Println(students)
	fmt.Println(students[0])

	students[0] = 100

	fmt.Println(students[0])
	fmt.Printf("Lenght of students %v\n", len(students))

	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	// Arrays are deep copied by default in Go
	b := matrix
	b[0][0] = 100
	fmt.Println(matrix)
	fmt.Println(b)

	// Slices don't have set size

	mySlice := []int{1, 2, 3, 4}
	mySlice = append(mySlice, 5)
	fmt.Printf("\nCapacity of the slice : %v, Length of the slice: %v\n", cap(mySlice), len(mySlice))

	// Slices are not deep copied

	mySlice2 := mySlice
	mySlice2[0] = 100
	fmt.Println(mySlice)
	fmt.Println(mySlice2)
	fmt.Println(append(mySlice, mySlice2...))
}
