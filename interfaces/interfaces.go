package main

import "fmt"

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	fmt.Println(cw)

	return 1, nil
}

func main() {
	var w Writer = ConsoleWriter{}
	fmt.Println(w)
}
