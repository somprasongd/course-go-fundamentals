package main

import "fmt"

func main() {
	// จะมีค่าเป็น zero value คือ nil ยังใช้งานไม่ได้
	var nums []int

	// ต้อง allocate memory ให้ก่อนถึงจะใช้งานได้
	nums = make([]int, 4) // จะต้องระบุขนาดเริ่มต้นให้ก่อน ทุกตำแหน่งจะได้ค่า zero value ของ type นั้นๆ
	// []int{0, 0, 0, 0}

	// หรือสร้างแบบว่างๆ
	nums = []int{}

	// หรือ จะใส่ค่าไปตั้งแต่ประกาศเลยก็ได้
	nums = []int{1, 2, 3}

	// append() - เพิ่มค่าลงใน slice
	nums = append(nums, 4)

	fmt.Printf("%#v\n", nums) // []int{1, 2, 3, 4}

	// len() กับ cap() - ดูขนาด และดูพื้นที่สำหรับเก็บข้อมูลใน slice ตามลำดับ
	fmt.Println(len(nums)) // 4
	fmt.Println(cap(nums)) // 6

	// [index_เริ่มต้น:index_ที่ต้องการ+1] - สำหรับการ slice ข้อมูล
	//          0   1   2   3   4   5   6   7   8
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	// เอาตั้งแต่ 0 จนถึงตัวสุดท้าย
	y := x[0:]

	// หรือเขียนแบบนี้ก็ได้
	y = x[:]

	// ถ้าต้องการ 30, 40, 50
	y = x[2:5] // ค่า 30 คือ index ที่่ 2 ส่วน 50 คือ index ที่ 4 ดังนั้นต้อง + 1 = 5

	// ถ้าต้องการต้องแต่เริ่มต้น จนถึงตำแหน่งที่ต้องการ สามารถละ index เริ่มต้นได้
	y = x[:5]

	fmt.Println(y)

	// การลบข้อมูลออกจาก slice
	words := []string{"A", "B", "C", "D", "E"}

	// ถ้าต้องการจะลบ "C" ออกไป ซึ่งก็คือตำแหน่งที่ 2 -> words[:2]
	// จะต้องได้ {"A", "B"} + {"D", "E"}
	// แต่ไม่สามารถใส่ words[3:] ลงไปตรงๆ ได้
	// จะต้องใช้ spread operator แทนแบบนี้
	words = append(words[:2], words[3:]...)

	fmt.Println(words) // ผลลัพธ์: [A B D E]

	// การวนลูปด้วย for range
	nums = []int{1, 2, 3}
	for i, v := range nums {
		fmt.Println(i, v) // i คือ index, v คือ value
	}
}
