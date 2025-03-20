package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	// ตรวหาค่าใน context ด้วย Key "key"
	if value := ctx.Value("key"); value != nil {
		fmt.Println("Received value:", value) // ถ้ามีค่า, แสดงผลลัพธ์
	} else {
		fmt.Println("No value found")
	}
}

func main() {
	// สร้าง context ที่มีค่า "key" คือ "myValue"
	ctx := context.WithValue(context.Background(), "key", "myValue")
	process(ctx) // ส่งค่า "myValue" ผ่าน context
}

// ผลลัพธ์:
// Received value: myValue
