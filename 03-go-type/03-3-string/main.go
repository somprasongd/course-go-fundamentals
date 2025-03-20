package main

import "fmt"

func main() {
	str := "Hello, Go!"
	fmt.Println(str[0])         // 72 (ASCII ของ 'H')
	fmt.Println(string(str[0])) // H
}
