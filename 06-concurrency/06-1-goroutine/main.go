package main

import (
	"fmt"
	"time"
)

// ฟังก์ชันที่จะรันใน Goroutine
func worker(id int) {
	fmt.Printf("Worker %d started\n", id)
}

func main() {
	for i := 1; i <= 3; i++ {
		go worker(i) // เรียกใช้ฟังก์ชันใน Goroutine
	}

	fmt.Println("main function running")

	time.Sleep(500 * time.Millisecond)
	fmt.Println("All workers done!")
}

// ผลลัพธ์ (อาจเปลี่ยนแปลงได้ในแต่ละครั้ง):
// main function running
// Worker 3 started
// Worker 2 started
// Worker 1 started
// All workers done!
