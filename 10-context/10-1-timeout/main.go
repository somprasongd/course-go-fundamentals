package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second): // จำลองงานที่ใช้เวลานาน
		fmt.Println("Task completed")
	case <-ctx.Done(): // เมื่อเวลา timeout หรือมีการยกเลิก
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

func main() {
	// ใช้ context.WithTimeout เพื่อกำหนดเวลาสำหรับการทำงาน
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // ให้แน่ใจว่า cancel ถูกเรียกหลังจากการใช้งาน

	longRunningTask(ctx) // เรียกใช้งานฟังก์ชันที่ต้องรอผล
}

// ผลลัพธ์:
// Task cancelled: context deadline exceeded
