package main

import (
	"fmt"
)

func a() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover in a:", r)
		}
	}()
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
// Recover in a: Panic in b
// Completed
