package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fn := add             // เก็บฟังก์ชันในตัวแปร
	fmt.Println(fn(3, 4)) // เรียกใช้ฟังก์ชันผ่านตัวแปร
}

// ผลลัพธ์:
// 7
