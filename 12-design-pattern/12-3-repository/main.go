package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"` // กรณีที่เป็น 0 จะไม่แสดงใน JSON
	Address string `json:"-"`             // ไม่ให้แปลง Address ไปเป็น JSON
}

func main() {
	person := Person{
		Name:    "John Doe",
		Address: "123 Main St",
	}

	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))
}

// ผลลัพธ์:
// {"name":"John Doe"}
