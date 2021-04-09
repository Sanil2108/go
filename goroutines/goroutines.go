package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

var m = sync.RWMutex{}
var c = 0

func main() {
	wg.Add(2)
	go sayHello()
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}("Saying hello again")

	for i := 0; i < 5; i += 1 {
		wg.Add(2)
		m.RLock()
		go printCounter()
		m.Lock()
		go increment()
	}

	wg.Wait()

}

func increment() {
	c += 1
	m.Unlock()
	wg.Done()
}

func printCounter() {
	fmt.Println(c)
	m.RUnlock()
	wg.Done()
}

func sayHello() {
	fmt.Println("Hello!")
	wg.Done()
}
