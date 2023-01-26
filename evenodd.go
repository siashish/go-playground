// even odd using go routine
package main

import (
	"fmt"
	"sync"
)

func main() {
	var w sync.WaitGroup

	w.Add(3)
	go Even(&w)
	go Odd(&w)
	go Odd(&w)
	w.Wait()
	fmt.Println("main finished")

}

func Even(w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("even")
}

func Odd(w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Odd")
}
