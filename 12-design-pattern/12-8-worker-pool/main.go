package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker function
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d กำลังประมวลผลงาน %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // จำลองเวลาประมวลผล
		fmt.Printf("Worker %d จบการทำงาน %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	numWorkers := 3 // จำนวน worker ที่ต้องการใช้
	numJobs := 5    // จำนวนงานที่ต้องประมวลผล

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// สร้าง workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// ส่งงานเข้า queue
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // ปิด channel เพื่อบอก worker ว่าไม่มีงานใหม่แล้ว

	// รอให้ worker ทำงานเสร็จทั้งหมด
	wg.Wait()
	close(results)

	// แสดงผลลัพธ์
	for result := range results {
		fmt.Println("ผลลัพธ์:", result)
	}
}
