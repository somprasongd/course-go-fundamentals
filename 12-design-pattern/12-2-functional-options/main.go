package main

import (
	"encoding/json"
	"fmt"
)

// สร้าง struct ที่ใช้ในการแปลง JSON
type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	jsonData := `{"Name":"John Doe","Age":30,"Address":"123 Main St"}`

	var person Person

	// แปลง JSON ไปเป็น struct
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	// แสดงผล struct
	fmt.Println(person)
}

// ผลลัพธ์:
// {John Doe 30 123 Main St}
