package main

import "fmt"

func main() {
	fmt.Printf(`- แสดงข้อความใช้ %s
- แสดงตัวเลขใช้ %d
- แสดงค่าของตัวแปรใช้ %v
- แสดงชนิดของตัวแปรใช้ %T\n`,
		"ข้อความ", // %s: แสดงข้อความ
		123,       // %d: แสดงตัวเลข
		"ใส่อะไรมาก็ได้", // %v: แสดงค่าของตัวแปร
		true) // %T: แสดงชนิดของตัวแปร
}
