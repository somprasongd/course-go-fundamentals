package main

import (
	"fmt"
)

// ใช้ interface{} รองรับได้ทุกประเภท แต่ต้องแปลงค่ากลับ
func Sum(nums []interface{}) interface{} {
	var totalInt int
	var totalFloat float64

	for _, num := range nums {
		switch v := num.(type) {
		case int:
			totalInt += v
		case float64:
			totalFloat += v
		default:
			return fmt.Errorf("unsupported type: %T", v)
		}
	}

	if totalFloat != 0 {
		return totalFloat
	}
	return totalInt
}

func main() {
	fmt.Println(Sum([]interface{}{1, 2, 3}))       // 6
	fmt.Println(Sum([]interface{}{1.5, 2.5}))      // 4.0
	fmt.Println(Sum([]interface{}{"hello", 2, 3})) // error
}
