package main

import (
	"fmt"
)

func a() {
	b()
}

func b() {
	panic("Panic in b")
}

func main() {
	a()
	fmt.Println("Completed")
}

// ผลลัพธ์:
// panic: Panic in b
