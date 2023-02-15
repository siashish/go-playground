package main

import "fmt"

func exc(list1, list2 []byte) []byte {

	result := make([]byte, len(list1))
	for i := 0; i < len(list1); i++ {
		flag := false
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				flag = true
			}
		}
		if flag == false {
			result[i] = list1[i]
		}
	}
	return result
}

func main() {
	list1 := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N'}
	list2 := []byte{'B', 'E'}
	fmt.Println(string(list1), string(list2))
	fmt.Println(string(exc(list1, list2)))
}
