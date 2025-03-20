package main

import "fmt"

type Box[T any] struct {
	Value T
}

func main() {
	intBox := Box[int]{Value: 100}
	strBox := Box[string]{Value: "Hello"}

	fmt.Println(intBox.Value) // 100
	fmt.Println(strBox.Value) // Hello
}
