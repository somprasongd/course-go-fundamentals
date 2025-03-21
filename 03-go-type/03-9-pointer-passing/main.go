package main

import (
	"fmt"
)

func changeValue(x int) {
	x = 100 // เปลี่ยนค่า x ภายในฟังก์ชัน
	fmt.Println("ค่าภายในฟังก์ชัน:", x)
}

func changeValuePointer(x *int) {
	*x = 100 // ใช้ pointer เพื่อเปลี่ยนค่าของตัวแปรต้นฉบับ
	fmt.Println("ค่าภายในฟังก์ชัน [pointer]:", *x)
}

func main() {
	a := 10

	changeValue(a)
	fmt.Println("ค่าภายนอกฟังก์ชัน:", a) // ค่า a ไม่เปลี่ยนแปลง

	changeValuePointer(&a)                         // ส่ง pointer ของ a ไป
	fmt.Println("ค่าภายนอกฟังก์ชัน [pointer]:", a) // ค่า a ถูกเปลี่ยนแปลง
}

// ผลลัพธ์:
// ค่าภายในฟังก์ชัน: 100
// ค่าภายนอกฟังก์ชัน: 10
// ค่าภายในฟังก์ชัน [pointer]: 100
// ค่าภายนอกฟังก์ชัน [pointer]: 100
