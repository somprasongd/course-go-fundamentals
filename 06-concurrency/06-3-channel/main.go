package main

import "fmt"

func sendData(ch chan string) {
	ch <- "Hello from Goroutine" // ส่งข้อมูลผ่าน Channel
}

func main() {
	ch := make(chan string) // สร้าง Channel
	go sendData(ch)         // เรียก Goroutine
	msg := <-ch             // รอรับข้อมูลจาก Channel
	fmt.Println(msg)
}

// ผลลัพธ์:
// Hello from Goroutine
