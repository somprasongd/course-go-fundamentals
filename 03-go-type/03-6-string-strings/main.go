package main

import (
	"fmt"
	"strings"
)

func main() {
	// ตรวจสอบว่ามีคำนี้ในข้อความหรือไม่ case sensitive
	result1 := strings.Contains("Hello Gopher", "go")
	fmt.Println(result1) // false

	// นับคำที่ต้องการหาว่ามีกี่คำในข้อความ
	result2 := strings.Count("สวัสดีชาวโก", "ดี")
	fmt.Println(result2) // 1

	// ตรวจสอบคำขึ้นต้น
	result3 := strings.HasPrefix("สวัสดีชาวโก", "สวั")
	fmt.Println(result3) // true

	// ตรวจสอบคำลงท้าย
	result4 := strings.HasSuffix("สวัสดีชาวโก", "โก")
	fmt.Println(result4) // true

	// ต่อข้อความจาก []string
	result5 := strings.Join([]string{"สวัสดี", "ชาวโก"}, "_")
	fmt.Println(result5) // สวัสดี_ชาวโก

	// แปลงข้อความเป็นตัวพิมพ์ใหญ่ทั้งหมด
	result6 := strings.ToUpper("Hello Gopher")
	fmt.Println(result6) // HELLO GOPHER

	// แปลงข้อความเป็นตัวพิมพ์เล็กทั้งหมด
	result7 := strings.ToLower("Hello Gopher")
	fmt.Println(result7) // hello gopher
}
