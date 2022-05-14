package main

import "fmt"

func main() {
	var n, min int
	fmt.Scanf("%d", &n)

	len := n*2 - 1
	// Complete the code to print the pattern.
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if i < j {
				min = i
			} else {
				min = j
			}

			if min > len-i {
				min = len - i - 1
			}

			if min > len-j-1 {
				min = len - j - 1
			}
			fmt.Printf("%d ", n-min)
		}
		fmt.Println()
	}
}
