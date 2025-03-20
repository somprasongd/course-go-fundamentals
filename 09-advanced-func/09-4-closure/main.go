package main

import "fmt"

// ฟังก์ชันที่คืนค่าฟังก์ชัน (Closure)
func counter() func() int {
	count := 0 // ตัวแปรนี้อยู่ใน Scope ของ counter()
	return func() int {
		count++ // เพิ่มค่า count ทุกครั้งที่เรียกใช้งาน
		return count
	}
}

func main() {
	c1 := counter()   // สร้าง Closure ตัวแรก
	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2

	c2 := counter()   // สร้าง Closure ตัวใหม่ (count เริ่มต้นใหม่)
	fmt.Println(c2()) // 1
}

// ผลลัพธ์:
// 1
// 2
// 1
