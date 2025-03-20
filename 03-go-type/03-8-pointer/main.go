package main

import (
	"fmt"
)

func main() {
	name := "Go"

	var s *string // สร้าง pointer ของ string

	s = &name // ดึงค่า address ออกมา

	fmt.Println(*s) // ผลลัพท์: Go เนื่องจาก *s เข้าถึงค่าใน address ที่อ้างอิงถึง

	*s = "Golang"     // แก้ไขค่าใน address ที่อ้างอิงถึง
	fmt.Println(name) // Golang
}
