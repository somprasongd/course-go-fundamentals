package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // งานที่ใช้เวลานาน
		fmt.Println("Task completed")
	case <-ctx.Done(): // ถ้ามีการยกเลิก
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go longRunningTask(ctx) // เรียกใช้งานงานที่ต้องใช้เวลานาน

	time.Sleep(2 * time.Second)        // จำลองการทำงานในบางช่วงเวลา
	cancel()                           // ยกเลิกการทำงานหลังจาก 2 วินาที
	time.Sleep(500 * time.Millisecond) // หยุดเพื่อให้เห็นผลลัพธ์ถูกยกเลิก
}

// ผลลัพธ์:
// Task cancelled: context canceled
