package main

import "fmt"

func shortDeclaration() {
	// เอา var ออก แล้วใช้ := แทน =
	i := 20
	s := "hello"
	ok := true

	fmt.Println(i, s, ok)
}
func main() {
	shortDeclaration()
}
