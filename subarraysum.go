package main

import "fmt"

func subarray(nums []int, k int) int {
	count, sum := 0, 0
	Map := make(map[int]int)
	Map[0] = 1
	for _, n := range nums {
		sum += n
		count += Map[sum-k]
		Map[sum]++
	}
	return count
}

func main() {
	arr := []int{1, 2, 3, 6, -6, -2, -1, 4}
	count := subarray(arr, 0)
	fmt.Println(count)
}
