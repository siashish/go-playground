package main

import "fmt"

type Helllo interface {
	Say()
}

func Say(a string) string {
	return "hello " + a
}

func main() {
	fmt.Println(Say("Ashish"))
}
