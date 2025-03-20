package main

import (
	"fmt"
	"sync"
)

// ฟังก์ชันที่จะรันใน Goroutine
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // บอก WaitGroup ว่า Goroutine เสร็จแล้ว
	fmt.Printf("Worker %d started\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // เพิ่มจำนวน Goroutine
		go worker(i, &wg)
	}

	fmt.Println("main function running")

	wg.Wait() // รอจน Goroutine ทั้งหมดเสร็จ
	fmt.Println("All workers done!")
}

// ผลลัพธ์ (อาจเปลี่ยนแปลงได้ในแต่ละครั้ง):
// main function running
// Worker 3 started
// Worker 2 started
// Worker 1 started
// All workers done!
