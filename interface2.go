package main

import (
	"fmt"
	"math/cmplx"
)

func TypeCheck(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("Type: int Value: ", a.(int))
	case float64:
		fmt.Println("Type: float64 Value: ", a.(float64))
	case string:
		fmt.Println("Type: string Value: ", a.(string))
	case bool:
		fmt.Println("Type: bool Value: ", a.(bool))
	case complex128:
		fmt.Println("Type: complex128 Value: ", a.(complex128s))
	default:
		fmt.Println("Type: Unknown Value: ", a)
	}
}

func main() {
	// inside main function
	TypeCheck("Hello Folks! Welcome to Knoldus blogs")
	TypeCheck(67)
	TypeCheck(true)
	TypeCheck(12.23)
	TypeCheck(cmplx.Sqrt(-5 + 12i))
}
