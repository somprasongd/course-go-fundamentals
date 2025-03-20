package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// normal fucntion
func Area(rec Rectangle) float64 {
	return rec.Width * rec.Height
}

// recevier function
func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

func main() {
	rec := Rectangle{
		Width:  10,
		Height: 20,
	}

	// fucntion style
	fmt.Println(Area(rec))

	// method style
	fmt.Println(rec.Area())
}
