package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(10, 20, 30, 40))
}

// ผลลัพธ์:
// 6
// 100
