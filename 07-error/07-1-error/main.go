package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("5s")
	if err != nil {
		// อาจจะ log หรือส่งออกไปให้ที่อื่นจัดการต่อ
		fmt.Println("Error converting string to integer:", err)
		return
	}
	fmt.Println("Number of seconds:", n)
}

// ผลลัพธ์:
// Error converting string to integer: strconv.Atoi: parsing "5s": invalid syntax
