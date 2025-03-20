package main

import "fmt"

// สร้าง interface ที่กำหนด method Speak
type Speaker interface {
	Speak() string
}

// สร้าง struct ที่ implement interface Speaker
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

// ฟังก์ชันที่รับ argument เป็น interface Speaker
func introduce(speaker Speaker) {
	fmt.Println(speaker.Speak())
}

func main() {
	person := Person{Name: "Alice"}
	introduce(person) // จะใช้ method Speak ที่อยู่ใน struct Person
}

// ผลลัพธ์:
// Hello, my name is Alice
