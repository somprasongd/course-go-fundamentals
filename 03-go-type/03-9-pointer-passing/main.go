package main

import (
	"fmt"
)

func inc(n int) {
	n = n + 1
}

func incPointer(n *int) {
	*n = *n + 1
}

func main() {
	i := 1

	inc(i)
	fmt.Println(i) // ผลลัพท์: 1

	incPointer(&i)
	fmt.Println(i) // ผลลัพท์: 2
}
