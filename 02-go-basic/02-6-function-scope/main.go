package main

import "fmt"

func sayHello() {
	var message = "Hello, Go!"
	fmt.Println(message) // ✅ ใช้งานได้ภายในฟังก์ชัน
}

func main() {
	sayHello()
	// fmt.Println(message) ❌ Error: ใช้ตัวแปร message ที่อยู่นอก scope ไม่ได้
}
