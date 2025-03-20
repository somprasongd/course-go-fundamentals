package main

import (
	"fmt"
)

func readFile(filename string) error {
	return fmt.Errorf("ไม่พบไฟล์: %s", filename)
}

func main() {
	err := readFile("data.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// ผลลัพธ์:
// Error: ไม่พบไฟล์: data.txt
