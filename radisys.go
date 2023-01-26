package main

import "fmt"

func subarray(num []int, k int) int {
	count, sum := 0, 0
	Map := make(map[int]int)
	Map[0] = 1
	for _, n := range num {
		sum += n
		count += Map[sum-k]
		Map[sum]++
	}
	return count
}

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("end")
			return
		}
	}
}

func main() {
	arr := []int{2, 3, 5, 7, 8}
	count := subarray(arr, 14)
	fmt.Println(count)

	fmt.Println("=============")

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib(c, quit)
}
