package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string, msg string, delay time.Duration) {
	time.Sleep(delay)
	ch <- msg
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendData(ch1, "From Channel 1", 2*time.Second)
	go sendData(ch2, "From Channel 2", 1*time.Second)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}
}

// ผลลัพธ์:
// From Channel 2
// เนื่องจาก ch2 ใช้เวลา 1 วินาที แต่ ch1 ใช้เวลา 2 วินาที → select เลือก ch2 ก่อน
