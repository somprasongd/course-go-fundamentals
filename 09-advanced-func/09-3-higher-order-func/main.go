package main

import "fmt"

// Higher-Order Function ที่รับฟังก์ชันเป็นพารามิเตอร์
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(operate(3, 4, add))      // 7
	fmt.Println(operate(3, 4, multiply)) // 12
}

// ผลลัพธ์:
// 7
// 12
