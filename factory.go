package main

import "fmt"

func Print(data interface{}) {
	fmt.Println(data)
}

func Add(x, y int) {
	Print(x + y)
}

func main() {
	Print(10)
	Print("ashish")
	Add(10, 20)
}
