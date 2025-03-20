package main

import "fmt"

// ฟังก์ชันที่ใช้รับค่าเป็น Empty Interface (interface{})
func printAnything(value interface{}) {
	fmt.Println(value)
}

func main() {
	printAnything("Hello, World!") // รับค่า string
	printAnything(123)             // รับค่า int
	printAnything(true)            // รับค่า bool
}

// ผลลัพธ์:
// Hello, World!
// 123
// true
