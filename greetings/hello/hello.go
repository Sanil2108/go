package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, error := greetings.Hello("Sanil")
	// if error != nil {
	// 	log.Fatal(error)
	// }

	messages, error := greetings.Hellos([]string{"Sanil", "Koshka", "Prerna"})
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(messages)
}
