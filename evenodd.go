// even odd using go routine
package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

func main() {

	w.Add(2)
	go Even()
	go Odd()
	w.Wait()
	fmt.Println("main finished")

}

func Even() {
	defer w.Done()
	fmt.Println("even")
}

func Odd() {
	defer w.Done()
	fmt.Println("Odd")
}
