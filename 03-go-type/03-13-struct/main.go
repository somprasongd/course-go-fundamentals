package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func main() {
	rec := Rectangle{} // จะมีค่าเป็น empty -> {}

	// หรือจะกำหนดค่าไปเลยก็ได้

	rec = Rectangle{
		Width:  10,
		Height: 20,
	}

	// การอ่านค่า
	fmt.Println(rec.Width) // 10

	// การกำหนด/แก้ไขค่า
	rec.Width = 30
	fmt.Println(rec.Width) // 30
}
