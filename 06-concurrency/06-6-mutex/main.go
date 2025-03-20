package main

import (
	"fmt"
	"sync"
)

var (
	count = 0
	mu    sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock() // ล็อกก่อนเข้าถึงข้อมูล
		count++
		mu.Unlock() // ปลดล็อกเมื่อเสร็จ
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait() // รอให้ทุก goroutine เสร็จทำงาน
	fmt.Println("Final count:", count)
}

// ผลลัพธ์:
// Final count: 10000
