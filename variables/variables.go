package main

import (
	"fmt"
	"strconv"
)

var outsideVariable int = 100

func main() {
	// Declare a variable
	var i int
	i = 42

	j := 100

	fmt.Printf("%v, %T", i, i)
	fmt.Printf("\n%v, %T", j, j)
	fmt.Printf("\n%v, %T", float32(j), float32(j))
	// This converts to the ASCII value on that int
	fmt.Printf("\n%v, %T", string(j), string(j))
	fmt.Printf("\n%v, %T", strconv.Itoa(j), strconv.Itoa(j))

}
