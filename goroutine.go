package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Print(ch <-chan int) {
	defer wg.Done()
	fmt.Println(<-ch)
}
func Set(ch chan<- int, v int) {
	defer wg.Done()
	ch <- v
}
func main() {
	ch := make(chan int)
	fmt.Println("Go routine start")
	wg.Add(2)
	go Set(ch, 10)
	go Print(ch)
	wg.Wait()
	fmt.Println("Go routine End")
}
