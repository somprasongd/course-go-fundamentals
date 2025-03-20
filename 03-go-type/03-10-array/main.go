package main

import "fmt"

func main() {
	var nums [4]int = [4]int{1, 2, 3} // กำหนด array ขนาด 4

	fmt.Println("Array:", nums) // ผลลัพท์: Array: [1 2 3 0]

	// การเข้าถึงและแก้ไขค่า
	fmt.Println(nums[0]) // ผลลัพท์: 1
	fmt.Println(nums[1]) // ผลลัพท์: 2
	fmt.Println(nums[2]) // ผลลัพท์: 3
	fmt.Println(nums[3]) // ผลลัพท์: 0

	nums[3] = 4                 // แก้ไขค่าในตำแหน่งที่ 3
	fmt.Println("Array:", nums) // ผลลัพท์: Array: [1 2 3 4]

	// การวนลูปด้วย for range
	for i, v := range nums {
		fmt.Println(i, v) // i คือ index, v คือ value
	}
}
