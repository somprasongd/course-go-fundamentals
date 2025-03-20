package main

import (
	"fmt"
)

func main() {
	fmt.Println(1)
	fmt.Println(2)
	panic("Fail")
	fmt.Println(3) // ไม่ได้ถูกเรียกใช้
	fmt.Println(4) // ไม่ได้ถูกเรียกใช้
	fmt.Println(5) // ไม่ได้ถูกเรียกใช้
}

// ผลลัพธ์:
// 1
// 2
// panic: Fail
