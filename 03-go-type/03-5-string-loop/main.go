package main

import "fmt"

func main() {
	message := "Hello, สวัสดี"

	for i, char := range message {
		fmt.Printf("Index: %d, Code: %d, Char: %c\n", i, char, char)
	}
}

// ผลลัพธ์:
// Index: 0, Code: 72, Char: H
// Index: 1, Code: 101, Char: e
// Index: 2, Code: 108, Char: l
// Index: 3, Code: 108, Char: l
// Index: 4, Code: 111, Char: o
// Index: 5, Code: 44, Char: ,
// Index: 6, Code: 32, Char:
// Index: 7, Code: 3626, Char: ส
// Index: 10, Code: 3623, Char: ว
// Index: 13, Code: 3633, Char: ั
// Index: 16, Code: 3626, Char: ส
// Index: 19, Code: 3604, Char: ด
// Index: 22, Code: 3637, Char: ี
