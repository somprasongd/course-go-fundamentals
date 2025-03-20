package main

import "fmt"

type Shape struct {
	Name string
}

type Rectangle struct {
	Width  float64
	Height float64
	Shape  // Embedding Shape
}

func main() {
	rec := Rectangle{
		Width:  10,
		Height: 20,
		Shape:  Shape{Name: "Rectangle"},
	}

	fmt.Println(rec.Shape.Name) // ผลลัพธ์: Rectangle
	// คุณสมบัติการ promoted fields
	fmt.Println(rec.Name) // ผลลัพธ์: Rectangle
}
