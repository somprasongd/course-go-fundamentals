package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Value receiver
func (r Rectangle) setWidth(w float64) {
	r.Width = w // การเปลี่ยนแปลงนี้จะไม่ส่งผลกับตัวแปร rect ที่ถูกเรียก
}

// Pointer receiver
func (r *Rectangle) setWidthPtr(w float64) {
	r.Width = w // การเปลี่ยนแปลงนี้จะส่งผลกับตัวแปร rect ที่ถูกเรียก
}

func main() {
	rect := Rectangle{Width: 10, Height: 20}
	rect.setWidth(20)
	fmt.Println("Width after set:", rect.Width) // ผลลัพธ์: 10

	rect.setWidthPtr(20)
	fmt.Println("Width after set with pointer:", rect.Width) // ผลลัพธ์: 20
}
