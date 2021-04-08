package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getRobots() {
	res, err := http.Get("http://google.com/robots.txt")
	if err != nil {
		log.Fatal("Something bad happened")
	}

	fmt.Println(res.Body)
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal("Something went wrong but something else")
	}

	fmt.Println(string(robots))
}

func main() {
	getRobots()

	// fmt.Println("start")
	// // All deferred functions are called after the rest of the function is over but has not returned
	// // Deferred calls runs LIFO
	// defer fmt.Println("middle")
	// fmt.Println("end")

	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

}
