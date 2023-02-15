// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func reversestring(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
		result = strings.ToUpper(result)
	}
	return
}

func main() {
	str := "hello"
	fmt.Println(str)
	fmt.Println(reversestring(str))
}
