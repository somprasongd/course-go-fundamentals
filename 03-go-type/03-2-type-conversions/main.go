package main

import "fmt"

func main() {
	var age int = 30
	var price float64 = float64(age) // แปลง int เป็น float64

	fmt.Println("Age as float64:", price)

	var num float64 = 42.5
	var intNum int = int(num) // แปลง float64 เป็น int (ปัดเศษทิ้ง)

	fmt.Println("Float to int:", intNum)
}
