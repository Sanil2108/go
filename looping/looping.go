package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i += 1 {
		fmt.Println(i)
	}

	i := 1
	j := 2

	i, j = j, i
	fmt.Println(i, j)

	for j < 10 {
		fmt.Println(1 / float32(j))
		j++
	}

	for k := 0; k < 10; k += 1 {
		for u := k; u < 10; u += 1 {
			fmt.Println(k * u)
		}
	}

	fmt.Println("\n\n")

	tempArray := []int{1, 2, 3}
	for k, v := range tempArray {
		fmt.Println(k, v)
	}

	tempMap := map[string]int{
		"Hello": 1,
		"World": 2,
	}
	for k, v := range tempMap {
		fmt.Println(k, v)
	}
}
