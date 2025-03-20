package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, สวัสดี"

	fmt.Println("ความยาวในหน่วย byte:", len(str)) // ผลลัพธ์: 17 (เนื่องจากอักษรภาษาไทยใช้หลาย byte ใน UTF-8)

	fmt.Println("จำนวน rune (ตัวอักษร):", utf8.RuneCountInString(str)) // ผลลัพธ์: 13

}
