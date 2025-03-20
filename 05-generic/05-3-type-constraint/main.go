package main

import "fmt"

// T คือ Type Parameter ที่รองรับ int และ float64
func Sum[T int | float64](nums []T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))      // 6
	fmt.Println(Sum([]float64{1.5, 2.5})) // 4.0
}
