package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a

	a = 44

	fmt.Println(a, *b)
}
