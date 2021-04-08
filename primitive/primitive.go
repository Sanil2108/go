package main

import (
	"fmt"
)

func main() {
	var boolValue bool = true
	fmt.Printf("%v, %T", boolValue, boolValue)

	var complex1 complex64 = 1 + 2i
	var complex2 complex64 = 4 + 12i

	fmt.Printf("\n\n%v, %T", complex1*complex2, complex1)
	fmt.Printf("\n\n%v, %T", complex1/complex2, complex1)
	fmt.Printf("\n\n%v, %T", real(complex1), real(complex1))
	fmt.Printf("\n\n%v, %T", complex(5, 128), complex(5, 128))

	var s string = "Hello World"
	fmt.Printf("\n\n%v %v %v, %T", s, s+" Sanil", []byte(s), s)

	var r rune = 'a'
	fmt.Printf("\n\n%v, %T", r, r)

}
