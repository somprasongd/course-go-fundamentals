package main

import "fmt"

func main() {
	ch := make(chan int, 3) // Buffered Channel ที่รองรับ 3 ค่า

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch) // ดึงค่าออก
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// ผลลัพธ์:
// 1
// 2
// 3
