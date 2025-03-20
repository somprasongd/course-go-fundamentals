package main

import (
	"fmt"
	"strconv"
)

func main() {
	// แปลงเป็น float ต้องระบุขนาดเสมอ (32/64)
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f) // 3.14

	// แปลงเป็น int ต้องระบุเลขฐานของข้อความที่จะแปลงด้วย
	i, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println(i) // -42

	// แปลงเป็น int ต้องระบุเลขฐานของข้อความที่จะแปลงด้วย
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println(u) // 42

	// แปลงเป็น bool
	b, _ := strconv.ParseBool("true")
	fmt.Println(b) // true

	// แปลงข้อความเป็นจัวเลข
	i2, _ := strconv.Atoi("-42")
	fmt.Println(i2) // -42

	// แปลงตัวเลขเป็นข้อความ
	s := strconv.Itoa(-42)
	fmt.Println(s) // "-42"
}
