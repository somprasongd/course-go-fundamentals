package main

import "fmt"

func greeting() {
	fmt.Println("Hello")
}

func greetingWithName(name string) {
	fmt.Println("Hello", name)
}

// สามารถประกาศชนิดของค่าที่คืนกลับได้หลังวงเล็บพารามิเตอร์
func sum(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func swap2(a, b int) (x int, y int) {
	x = b
	y = a
	return
}

func main() {
	greeting()
	greetingWithName("John")
	fmt.Println(sum(10, 20))
	fmt.Println(swap2(10, 20))
}
