package main

import "fmt"

var name = "GoLang"

func main() {
	fmt.Println(name) // "GoLang"

	name := "Gopher"
	fmt.Println(name) // "Gopher" (ใช้ตัวแปรใหม่ภายใน main)

	if true {
		name := "Inside Block"
		fmt.Println(name) // "Inside Block" (ตัวแปรใหม่ในบล็อก if)
	}

	fmt.Println(name) // "Gopher" (กลับมาใช้ตัวแปรของ main)
}
