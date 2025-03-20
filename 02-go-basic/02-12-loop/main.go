package main

import (
	"fmt"
)

func normal() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func while() {
	sum := 0
	i := 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

func infinite() {
	for {
		fmt.Println("Running...")
	}
}

func forRange() {
	nums := []int{1, 2, 3}

	for index, value := range nums {
		fmt.Println(index, value)
	}

	// ถ้าไม่ต้องการ index
	for _, value := range nums {
		fmt.Println(value)
	}

	// ถ้าต้องการเฉพาะ index
	for index := range nums {
		fmt.Println(index)
	}
}

func main() {
	normal()
	while()
	forRange()

	infinite()
}
