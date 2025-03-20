package main

import "fmt"

func checkType(value interface{}) {
	if str, ok := value.(string); ok {
		fmt.Println("This is a string:", str)
	} else if num, ok := value.(int); ok {
		fmt.Println("This is an integer:", num)
	} else {
		fmt.Println("Unknown type")
	}
}

func main() {
	checkType("Hello, Go!") // String
	checkType(42)           // Integer
	checkType(3.14)         // Unknown type
}

// ผลลัพธ์:
// This is a string: Hello, Go!
// This is an integer: 42
// Unknown type
