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
	case <-ctx.Done(): // เมื่อถึงเวลา deadline หรือมีการยกเลิก
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

func main() {
	// กำหนด deadline เป็นเวลาปัจจุบัน + 1 วินาที
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // ให้แน่ใจว่า cancel ถูกเรียกหลังจากการใช้งาน

	longRunningTask(ctx) // เรียกใช้งานฟังก์ชันที่ต้องรอผล
}

// ผลลัพธ์:
// Task cancelled: context deadline exceeded
