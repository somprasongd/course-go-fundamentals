package main

import (
	"encoding/json"
	"fmt"
)

// สร้าง struct ที่จะใช้ในการแปลง
type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
	}

	// แปลง struct ไปเป็น JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	// แสดงผล JSON ในรูปแบบ string
	fmt.Println(string(jsonData))
}

// ผลลัพธ์:
// {"Name":"John Doe","Age":30,"Address":"123 Main St"}
