package main

import (
	"fmt"
	"sync"
)

var (
	count = 0
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		count++ // อัปเดตค่าพร้อมกัน อาจเกิด Race Condition
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()                          // รอให้ทุก goroutine เสร็จทำงาน
	fmt.Println("Final count:", count) // ค่าผิดพลาดเพราะ Race Condition
}

// go run -race main.go
// ผลลัพธ์:
// Final count: 7051
// Found 3 data race(s)
