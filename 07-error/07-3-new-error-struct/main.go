package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return errors.New("เกิดข้อผิดพลาดบางอย่าง")
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// ผลลัพธ์:
// Error: เกิดข้อผิดพลาดบางอย่าง
