# Go Fundamentals

Source code for Go Fundamentals Course

Blog:

- [Go Fundamentals](https://somprasongd.work/blog/go/go-fundamentals) (Update: 2025-03-21)
- [Design Patterns ที่นิยมใช้ในภาษา Go](https://somprasongd.work/blog/go/desing-patterns-in-go) (Update: 2025-03-17)

---

ภาษา Go (หรือ Golang) เป็นภาษาโปรแกรมระดับสูงที่พัฒนาโดย Google โดยเน้นความเรียบง่าย ประสิทธิภาพสูง และรองรับการประมวลผลพร้อมกัน (Concurrency) ได้อย่างยอดเยี่ยม ด้วยความสามารถเหล่านี้ทำให้ Go ได้รับความนิยมอย่างรวดเร็วในหมู่นักพัฒนาซอฟต์แวร์ โดยเฉพาะอย่างยิ่งในระบบที่ต้องการความเร็วและความเสถียรสูง เช่น ระบบเครือข่าย บริการแบบกระจาย (Distributed Systems) และแอปพลิเคชันที่มีการประมวลผลพร้อมกัน

บทความนี้จะพาคุณไปรู้จักกับภาษา Go ตั้งแต่ขั้นพื้นฐานจนถึงแนวคิดขั้นสูง โดยครอบคลุมหัวข้อสำคัญ ได้แก่

- [เริ่มต้นใช้งาน](#1-เริ่มต้นใช้งาน)
- [พื้นฐานภาษา Go](#2-พื้นฐานภาษา-go)
- [ชนิดของข้อมูลใน Go](#3ชนิดของข้อมูลใน-go)
- [ออกแบบโค้ดให้ยืดหยุ่นด้วย Interface](#4-ออกแบบโค้ดให้ยืดหยุ่นด้วย-interface)
- [ทำความรู้จักกับ Generic ในภาษา Go](#5-ทำความรู้จักกับ-generic-ในภาษา-go)
- [การทำงานแบบ Concurrency](#6-การทำงานแบบ-concurrency)
- [การจัดการข้อผิดพลาด (Error Handling)](#7-การจัดการข้อผิดพลาด-error-handling)
- [Unit Testing](#8-unit-testing)
- [การใช้งานฟังก์ชันขั้นสูง](#9-การใช้งานฟังก์ชันขั้นสูง)
- [การใช้ Context ใน Go](#10-การใช้-context-ใน-go)
- [การทำงานกับ JSON](#11-การทำงานกับ-json)

---

## 1. เริ่มต้นใช้งาน

เรามาเริ่มจากการติดตั้งภาษา Go และลองการเขียนโปรแกรม "Hello World" กันก่อน

- [การเตรียมเครื่อง](#การเตรียมเครื่อง)
- [การเขียนโปรแกรม "Hello World"](#การเขียนโปรแกรม-hello-world)
- [การรันโปรแกรม](#การรันโปรแกรม)
- [การใช้งาน Go package](#การใช้งาน-go-package)

---

### การเตรียมเครื่อง

- ติดตั้ง Go จาก[ที่นี่](https://go.dev/dl/) (ปัจจุบันเวอร์ชัน 1.24.1)
- ติดตั้ง VS Code จาก[ที่นี่](https://code.visualstudio.com/download) และติดตั้ง extensions เพิ่มเติมดังนี้
  - [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) ต้องติดตั้ง
  - [Error Lens](https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens) แนะนำไว้เป็นตัวช่วยแสดงข้อผิดพลาด

---

### การเขียนโปรแกรม "Hello World"

การสร้างโปรเจกต์ใหม่ใน Go ให้เริ่มจากสร้าง **Go module** ก่อน เพื่อเปิดใช้งาน **dependency tracking** มีวิธีดังนี้:

1. **สร้าง Directory สำหรับโปรเจกต์**

   ```bash
   mkdir hello
   cd hello
   ```

2. **สร้าง Go Module**

   แนะนำให้ตั้งชื่อ module ตามชื่อ repository หากต้องการแชร์โค้ดให้ผู้อื่นใช้งาน

   ```bash
   go mod init gobasic
   ```

3. **โครงสร้างพื้นฐานของโปรแกรม Go**

   Go มีข้อกำหนดสำคัญ 2 ข้อสำหรับจุดเริ่มต้นของโปรแกรม

   - ต้องมี `package main`
   - โปรแกรมเริ่มทำงานที่ฟังก์ชัน `main`

   > ไฟล์นี้สามารถอยู่ที่ไหนและตั้งชื่ออะไรก็ได้ แต่โดยทั่วไปจะตั้งชื่อเป็น main.go

   **ตัวอย่างโค้ด "Hello World":**

   ```go
   package main

   import "fmt"

   func main() {
     fmt.Println("Hello world!") // พิมพ์ข้อความออกทาง console
   }
   ```

4. **การจัดรูปแบบการแสดงผล**

   Go รองรับการพิมพ์ข้อความแบบ format ด้วย `fmt.Printf()`

   **ตัวอย่าง:**

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Printf(`- แสดงข้อความใช้ %s
   - แสดงตัวเลขใช้ %d
   - แสดงค่าของตัวแปรใช้ %v
   - แสดงชนิดของตัวแปรใช้ %T\n`,
           "ข้อความ",     // %s: แสดงข้อความ
           123,          // %d: แสดงตัวเลข
           "ใส่อะไรมาก็ได้", // %v: แสดงค่าของตัวแปร
           true)         // %T: แสดงชนิดของตัวแปร
   }
   ```

---

### การรันโปรแกรม

ใน Go สามารถรันโปรแกรมได้ 2 วิธีหลัก ๆ คือ **รันโปรแกรมโดยตรง** และ **คอมไพล์เป็น Binary File**

1. **รันโปรแกรมโดยตรง**

   ใช้คำสั่ง `go run` เพื่อตรวจสอบและรันโค้ดทันที:

   ```bash
   go run main.go
   ```

   > 💡 เหมาะสำหรับการทดสอบหรือพัฒนา เพราะไม่สร้างไฟล์ executable

2. **คอมไพล์เป็น Binary File**

   หากต้องการนำโปรแกรมไปใช้งานหรือแจกจ่าย สามารถคอมไพล์ด้วย `go build`:

   ```bash
   go build main.go
   ```

   จากนั้นรันไฟล์ที่คอมไพล์แล้ว:

   ```bash
   ./main   # บน Linux/macOS
   main.exe # บน Windows
   ```

   > 💡 การใช้ `go build` จะสร้างไฟล์ executable ใน directory ปัจจุบัน

---

### การใช้งาน Go package

ใน Go เราไม่จำเป็นต้องเขียนโค้ดทั้งหมดในไฟล์เดียวกัน แต่สามารถแบ่งโค้ดออกเป็นส่วนๆ เพื่อให้โปรเจกต์ดูเป็นระเบียบและง่ายต่อการจัดการ

สิ่งที่ช่วยให้เราแยกโค้ดออกเป็นส่วนๆ ใน Go ก็คือ **package** (ซึ่งเราต้องสร้าง **Go Module** ก่อน) โดยเราสามารถแบ่ง **package** ตาม **โครงสร้างของ directory** ได้เลย

การทำแบบนี้มีขั้นตอนง่ายๆ ดังนี้:

1. **สร้าง Directory สำหรับ Package**

   ตั้งชื่อ directory ให้ตรงกับชื่อ package ที่ต้องการ เช่น หากต้องการสร้าง `package greet` ให้สร้าง directory ชื่อ `greet`:

   ```go
   mkdir greet
   ```

2. **สร้างไฟล์แรกของ Package**

   แนะนำให้ตั้งชื่อไฟล์แรกให้ตรงกับชื่อ package เพื่อความชัดเจน เช่น `greet.go`:

   ```bash
   cd greet
   touch greet.go
   ```

3. **เขียนโค้ดใน Package**

   ใน Go หากต้องการให้ **ตัวแปร (Variable)** หรือ **ฟังก์ชัน (Function)** สามารถเรียกใช้งานจาก package อื่นได้ ต้องตั้งชื่อขึ้นต้นด้วย **ตัวอักษรพิมพ์ใหญ่ (Exported)**:

   ```go
   // greet/greet.go
   package greet

   import "fmt"

   // ฟังก์ชันนี้สามารถถูกเรียกใช้จากภายนอกได้
   func Hi() {
       fmt.Println("Hi 👋")
   }
   ```

4. **การเรียกใช้งาน Package**

   ในไฟล์ `main.go` สามารถ import และเรียกใช้ฟังก์ชันจาก package ได้ดังนี้:

   ```go
   // main.go
   package main

   import "gobasic/greet"

   func main() {
       greet.Hi() // แสดงผล: Hi 👋
   }
   ```

   > 💡 **หมายเหตุ:** ชื่อที่ใช้ใน `import` ต้องตรงกับ path ของ package ที่กำหนดไว้ใน go module เช่น `gobasic/greet` หากอยู่ใน module ชื่อ `gobasic`

---

## 2. พื้นฐานภาษา Go

ทำความเข้าใจการประกาศตัวแปร ฟังก์ชัน และโครงสร้างควบคุมการทำงาน (Control Flow) ในภาษา Go

- [การประกาศตัวแปร (Variable Declaration)](#การประกาศตัวแปร-variable-declaration)
- [การทำงานของฟังก์ชัน (Function)](#การทำงานของฟังก์ชัน-function)
- [การควบคุมการทำงาน (Control Flow)](#การควบคุมการทำงาน-control-flow)

---

### การประกาศตัวแปร (Variable Declaration)

ในภาษา Go สามารถประกาศตัวแปรได้หลายรูปแบบ ดังนี้:

1. **การประกาศแบบกำหนดชนิดข้อมูล (Explicit Type)**

   ใช้ `var` ตามด้วยชื่อตัวแปรและชนิดของข้อมูล (Type)

   ```go
   // ถ้าไม่กำหนดค่า จะได้ค่าเริ่มต้น (Zero value)
   var i int     // 0
   var s string  // ""
   var ok bool   // false

   // กำหนดค่าตอนประกาศ
   var i int = 20
   var s string = "hello"
   var ok bool = true
   ```

2. **การประกาศแบบไม่กำหนดชนิดข้อมูล (Type Inference)**

   เมื่อกำหนดค่าให้ตัวแปร Go จะอนุมานชนิดข้อมูลให้อัตโนมัติ

   ```go
   // ถ้ากำหนดค่าแล้วละ type ได้
   var i = 20
   var s = "hello"
   var ok = true
   ```

3. **การประกาศแบบ Short Declaration (`:=`)**

   เป็นรูปแบบย่อสำหรับการประกาศตัวแปร ใช้ในกรณีที่อยู่ภายในฟังก์ชันเท่านั้น

   ```go
   // เอา var ออก แล้วใช้ := แทน =
   i := 20
   s := "hello"
   ok := true
   ```

   > 💡 Scope ของตัวแปรใน Go
   >
   > - **Package Scope**: ตัวแปรที่ประกาศนอกฟังก์ชัน ใช้งานได้ใน package เดียวกัน
   > - **Function Scope**: ตัวแปรที่ประกาศในฟังก์ชัน ใช้ได้เฉพาะในฟังก์ชันนั้น
   > - **Block Scope**: ตัวแปรที่ถูกประกาศ **ภายใน `{}`** จะสามารถใช้งานได้เฉพาะในบล็อกนั้นเท่านั้น

---

### ค่าคงที่ (Constant)

การประกาศค่าคงที่ (Constant) ใช้ `const` แทน `var` และต้องกำหนดค่าตั้งแต่ประกาศ

```go
const i int = 20
const s string = "hello"
const ok bool = true
```

**ประยุกต์ใช้สร้าง `enum` ร่วมกับ `iota`**

`iota` ใช้สร้างลำดับเลขอัตโนมัติ เริ่มจาก `0` และเพิ่มขึ้นทีละ `1`

```go
const (
    sunday = iota // 0
    monday        // 1
    tuesday       // 2
    wednesday     // 3
    thursday      // 4
    friday        // 5
    saturday      // 6
)
```

หากไม่ต้องการใช้ค่าบางตัว ให้ใช้ `_` แทน แต่ไม่สามารถข้ามลำดับได้

```go
const (
    a = iota // 0
    _        // ข้าม 1
    b        // 2
)
```

---

### การทำงานของฟังก์ชัน (Function)

ฟังก์ชันใน Go ใช้ keyword `func` ตามด้วยชื่อฟังก์ชัน พารามิเตอร์ (ถ้ามี) และค่าที่คืนกลับ (ถ้ามี)

รูปแบบฟังก์ชัน แบบต่างๆ:

1. ฟังก์ชันพื้นฐาน (Basic Function)

   ```go
   func greeting() {
       fmt.Println("Hello")
   }
   ```

2. ฟังก์ชันที่รับพารามิเตอร์ (Function with Parameters)

   ```go
   func greeting(name string) {
       fmt.Println("Hello", name)
   }
   ```

3. ฟังก์ชันที่คืนค่า (Function with Return Value)

   ```go
   // สามารถประกาศชนิดของค่าที่คืนกลับได้หลังวงเล็บพารามิเตอร์
   func sum(a, b int) int {
       return a + b
   }
   ```

4. ฟังก์ชันที่คืนค่าหลายตัว (Multiple Return Values)

   ```go
   func swap(a, b int) (int, int) {
       return b, a
   }
   ```

5. การใช้ Named Return

   สามารถตั้งชื่อค่าที่จะคืนกลับได้ โดยกำหนดไว้หลังวงเล็บพารามิเตอร์

   ```go
   func swap(a, b int) (x int, y int) {
       x = b
       y = a
       return
   }
   ```

   > 💡 การใช้ Named Return ช่วยให้โค้ดอ่านง่ายขึ้น โดยเฉพาะเมื่อมีการคืนค่าหลายตัว

---

### การควบคุมการทำงาน (Control Flow)

ในภาษา Go เราสามารถควบคุมลำดับการทำงานของโปรแกรมได้ด้วยโครงสร้างต่อไปนี้:

1. **เงื่อนไข (`if-else`, `switch`):** ใช้ตรวจสอบเงื่อนไขและกำหนดการทำงานตามค่าที่ได้

   - `if-else`: ใช้ตรวจสอบเงื่อนไขและเลือกทำงานตามค่าที่ได้

     ใน Go `if-else` มีความแตกต่างจากภาษาอื่น คือ ไม่ต้องใส่วงเล็บ `()` รอบเงื่อนไข:

     - **รูปแบบทั่วไป**

       ```go
       func pow(x, n, lim float64) float64 {
           v := math.Pow(x, n)

           if v < lim {
               return v
           } else {
               fmt.Printf("%g >= %g\n", v, lim)
           }
           return lim
       }
       ```

     - **แบบใช้ร่วมกับ Short Statement** (เหมาะสำหรับการประกาศตัวแปรที่ใช้เฉพาะในเงื่อนไข)

       ```go
       func pow(x, n, lim float64) float64 {
           if v := math.Pow(x, n); v < lim {
               return v
           }
           fmt.Printf("%g >= %g\n", v, lim) // ❌ error
           return lim
       }
       ```

       > 💡 ตัวแปรที่ประกาศใน Short Statement จะใช้ได้เฉพาะในบล็อก `if-else` เท่านั้น

   - `switch`: ใช้เลือกทำงานตามค่าของตัวแปร โดยไม่ต้องใช้ `if-else` หลายชั้น

     ใน Go `switch` มีความแตกต่างจากภาษาอื่น ดังนี้:

     - ไม่ต้องใส่วงเล็บ `()`
     - ไม่ต้องใช้ `break` เพราะ Go จะ `auto break` ให้อัตโนมัติ
     - ใช้ `fallthrough` หากต้องการให้ทำงานต่อใน case ถัดไป

     **ตัวอย่างการใช้งาน `switch`**

     - **รูปแบบทั่วไป**

       ```go
       os := runtime.GOOS

       switch os {
       case "darwin":
           fmt.Println("macOS")
       case "linux":
           fmt.Println("Linux")
       default:
           fmt.Printf("Unknown OS: %s\n", os)
       }
       ```

     - **แบบใช้ร่วมกับ Short Statement**

       ```go
       switch os := runtime.GOOS; os {
       case "darwin":
           fmt.Println("macOS")
       case "linux":
           fmt.Println("Linux")
       default:
           fmt.Printf("Unknown OS: %s\n", os)
       }
       ```

     - **แบบไม่มีเงื่อนไข (Switch True)**

       ```go
       os := runtime.GOOS

       switch {
       case os == "darwin":
           fmt.Println("macOS")
       case os == "linux":
           fmt.Println("Linux")
       default:
           fmt.Printf("Unknown OS: %s\n", os)
       }
       ```

       > 💡 ใช้ `switch true` เพื่อเขียนเงื่อนไขที่ซับซ้อนหลายแบบ

2. **การวนลูป (`for`):** ใช้ทำงานซ้ำ ๆ ตามเงื่อนไขที่กำหนด

   ในภาษา Go มีเพียง `for` เป็นคำสั่งสำหรับการวนลูป แต่สามารถเขียนได้หลายรูปแบบเพื่อให้เหมาะกับการใช้งานต่างๆ เช่น:

   - **รูปแบบมาตรฐาน** (เหมือน `for` ในภาษาอื่น ต่างตรงที่ไม่ต้องใส่ `()`)

     ```go
     sum := 0
     for i := 0; i < 10; i++ {
         sum += i
     }
     fmt.Println(sum)
     ```

   - **แบบ `while loop`** (ใส่เฉพาะเงื่อนไข)

     ```go
     sum := 0
     i := 0
     for i < 10 {
         sum += i
         i++
     }
     fmt.Println(sum)
     ```

   - **แบบ Infinite Loop** (ลูปไม่มีที่สิ้นสุด)

     ```go
     for {
         fmt.Println("Running...")
     }
     ```

   - **แบบ `for range`** (ใช้กับ Array, Slice, Map, Channel)

     ```go
     nums := []int{1, 2, 3}

     for index, value := range nums {
         fmt.Println(index, value)
     }

     // ถ้าไม่ต้องการ index
     for _, value := range nums {
         fmt.Println(value)
     }

     // ถ้าต้องการเฉพาะ index
     for index := range nums {
         fmt.Println(index)
     }
     ```

3. **การหน่วงการทำงาน (`defer`):** ใช้เลื่อนการทำงานของคำสั่งไปจนกว่าฟังก์ชันจะสิ้นสุด

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Start")

       defer fmt.Println("This will run at the end")

       fmt.Println("Doing something")
   }

   // ผลลัพธ์:
   // Start
   // Doing something
   // This will run at the end
   ```

   **⚠️ ข้อควรระวังเกี่ยวกับ `defer`:**

   - ค่าของ Arguments จะถูกประมวลผลทันทีเมื่อ `defer` ถูกเรียก
   - ลำดับการทำงานของ `defer` เป็นแบบ Stack (LIFO: Last In, First Out)

   ```go
   package main

   import "fmt"

   func main() {
    fmt.Println("counting")

    for i := 0; i < 3; i++ {
     defer fmt.Println(i)
       // เท่ากับแบบนี้
     // defer fmt.Println(0)
     // defer fmt.Println(1)
     // defer fmt.Println(2)
    }

    fmt.Println("done")
   }

   // ผลลัพธ์:
   // counting
   // done
   // 2
   // 1
   // 0
   ```

   `defer` เหมาะสำหรับการปิดไฟล์, ปิดการเชื่อมต่อ, หรือคืนทรัพยากร

   **ตัวอย่าง การปิดไฟล์ด้วย `defer`:**

   ```go
   package main

   import (
       "fmt"
       "os"
   )

   func main() {
       file, err := os.Open("example.txt")
       if err != nil {
           fmt.Println("Error opening file:", err)
           return
       }
       defer file.Close() // ✳️ file.Close() จะถูกเรียกใช้เมื่อฟังก์ชัน main() สิ้นสุดลง

       fmt.Println("File opened successfully")
   }
   ```

---

## 3. ชนิดของข้อมูลใน Go

ภาษา Go มีระบบชนิดข้อมูล (type) ที่ชัดเจนและเข้มงวด (strongly typed) เพื่อช่วยให้การจัดการข้อมูลมีประสิทธิภาพและลดข้อผิดพลาดในการเขียนโปรแกรม

- [ชนิดข้อมูลพื้นฐาน (Basic Types)](#ชนิดข้อมูลพื้นฐาน-basic-types)
- [การแปลงประเภท (Type Conversions)](#การแปลงประเภท-type-conversions)
- [การทำงานกับข้อความ (Strings)](#การทำงานกับข้อความ-strings)
- [ตัวชี้หน่วยความจำ (Pointer)](#ตัวชี้หน่วยความจำ-pointer)
- [ชนิดข้อมูลเชิงโครงสร้าง (Composite Types)](#ชนิดข้อมูลเชิงโครงสร้าง-composite-types)

---

### ชนิดข้อมูลพื้นฐาน (Basic Types)

ข้อมูลพื้นฐานที่ใช้บ่อย เช่น

| ชนิดข้อมูล | คำอธิบาย                 | Zero-Value | ตัวอย่าง                 |
| ---------- | ------------------------ | ---------- | ------------------------ |
| `int`      | จำนวนเต็ม                | `0`        | `int age = 30`           |
| `float64`  | จำนวนจริง (ทศนิยม)       | `0`        | `float64 price = 99.99`  |
| `string`   | ข้อความ                  | `""`       | `string name = "GoLang"` |
| `bool`     | ค่าความจริง (true/false) | `false`    | `bool isActive = true`   |

**ตัวอย่างการใช้งาน:**

```go
package main

import "fmt"

func main() {
    var age int = 30
    var price float64 = 99.99
    var name string = "GoLang"
    var isActive bool = true

    fmt.Println("Age:", age)
    fmt.Println("Price:", price)
    fmt.Println("Name:", name)
    fmt.Println("Is Active:", isActive)
}
```

---

### การแปลงประเภท (Type Conversions)

Go ไม่รองรับการแปลงประเภทโดยอัตโนมัติ (Implicit Conversion) เราต้องใช้ฟังก์ชันแปลงเอง (Explicit Conversion)

**ตัวอย่างการแปลงชนิดข้อมูล:**

```go
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
```

> ⚠️ การแปลงชนิดข้อมูลอาจทำให้ข้อมูลสูญหาย เช่น การแปลง `float64` เป็น `int` จะตัดทศนิยมออก

---

### การทำงานกับข้อความ (Strings)

ในภาษา Go ชนิดข้อมูล `string` ใช้สำหรับเก็บข้อความ และมีคุณสมบัติดังนี้:

1. เป็นลำดับของไบต์ (Bytes) หมายความว่าการเข้าถึงตัวอักษรแต่ละตัวจะได้ค่าตัวเลข ASCII หรือ UTF-8

   ```go
   s := "Hello"
   fmt.Println(s[0]) // 72 (ค่า ASCII ของ 'H')
   fmt.Println(string(str[0])) // H
   ```

2. เป็น immutable (เปลี่ยนแปลงค่าโดยตรงไม่ได้)

   ```go
   s := "Hello"
   s[0] = 'h' // ❌ Error! เปลี่ยนค่าโดยตรงไม่ได้
   ```

3. เก็บข้อมูลแบบ UTF-8 ทำให้รองรับหลายภาษา รวมถึงอักขระพิเศษ

   แต่ต้องระวังการเข้าถึงตัวอักษรเพราะอาจไม่ได้เป็นไบต์เดียวเสมอไป

   ```go
   s := "สวัสดี"
   fmt.Println(len(s)) // 18 (ไม่ใช่จำนวนตัวอักษร เพราะเป็นไบต์)
   ```

**การนับความยาวของข้อความ (String) ใน Go**

ในภาษา Go มี 2 วิธีหลักในการนับความยาวของข้อความ:

1. การนับจำนวนไบต์ด้วย `len()`

   ฟังก์ชัน `len()` ใช้เพื่อนับจำนวนไบต์ (`byte`) ที่ใช้เก็บ `string`

   ⚠️ **ข้อควรระวัง**: `len()` นับตามจำนวนไบต์ ไม่ใช่จำนวนตัวอักษร (rune) ดังนั้นอักษรที่ใช้หลายไบต์ (เช่น อักษรไทย, จีน, หรืออิโมจิ) อาจทำให้ค่าที่ได้ไม่ตรงกับจำนวนตัวอักษรจริง

   ```go

   package main

   import "fmt"

   func main() {
    s := "Hello"
    fmt.Println(len(s)) // ได้ค่า 5 (เพราะแต่ละตัวอักษรภาษาอังกฤษใช้ 1 ไบต์)

    s2 := "สวัสดี"
    fmt.Println(len(s2)) // ได้ค่า 18 (เพราะแต่ละตัวอักษรภาษาไทยใช้ 3 ไบต์)
   }
   ```

2. การนับจำนวนตัวอักษร (rune) ด้วย `utf8.RuneCountInString()`

   ฟังก์ชัน `utf8.RuneCountInString()` จากแพ็กเกจ `utf8` ใช้เพื่อนับจำนวนตัวอักษร (rune) หรืออักขระ Unicode ที่อยู่ใน `string`

   ```go
   package main

   import (
       "fmt"
       "unicode/utf8"
   )

   func main() {
       s := "สวัสดี"
       fmt.Println(utf8.RuneCountInString(s)) // ได้ค่า 6 (จำนวนตัวอักษรจริง)
   }
   ```

สรุปความแตกต่าง

| **วิธีการ**                 | **นับเป็น**          | **ใช้เมื่อ**                                           |
| --------------------------- | -------------------- | ------------------------------------------------------ |
| `len(s)`                    | จำนวนไบต์ (bytes)    | ต้องการวัดขนาดของข้อมูลในหน่วยความจำ                   |
| `utf8.RuneCountInString(s)` | จำนวนตัวอักษร (rune) | ต้องการนับจำนวนตัวอักษรที่ใช้จริง เช่น สำหรับการแสดงผล |

**การวนลูปใน String**

ใช้ `for range` เพื่อวนลูปผ่าน `string` โดยค่าที่ได้จะเป็น **`rune`** ซึ่งแทนตัวอักษร Unicode

**โค้ดตัวอย่าง:**

```go
package main

import (
 "fmt"
 "unicode/utf8"
)

func main() {
 s := "Hello, สวัสดี"

 for i, r := range s {
  fmt.Printf("Index: %d, \tRune: %c, \tASCII: %d, \tUnicode: %U, \tBytes: %d\n", i, r, r, r, utf8.RuneLen(r))
 }
}

// ผลลัพธ์:
// Index: 0,    Rune: H,    ASCII: 72,      Unicode: U+0048,    Bytes: 1
// Index: 1,    Rune: e,    ASCII: 101,     Unicode: U+0065,    Bytes: 1
// Index: 2,    Rune: l,    ASCII: 108,     Unicode: U+006C,    Bytes: 1
// Index: 3,    Rune: l,    ASCII: 108,     Unicode: U+006C,    Bytes: 1
// Index: 4,    Rune: o,    ASCII: 111,     Unicode: U+006F,    Bytes: 1
// Index: 5,    Rune: ,,    ASCII: 44,      Unicode: U+002C,    Bytes: 1
// Index: 6,    Rune:  ,    ASCII: 32,      Unicode: U+0020,    Bytes: 1
// Index: 7,    Rune: ส,    ASCII: 3626,    Unicode: U+0E2A,    Bytes: 3
// Index: 10,   Rune: ว,    ASCII: 3623,    Unicode: U+0E27,    Bytes: 3
// Index: 13,   Rune: ั,     ASCII: 3633,    Unicode: U+0E31,    Bytes: 3
// Index: 16,   Rune: ส,    ASCII: 3626,    Unicode: U+0E2A,    Bytes: 3
// Index: 19,   Rune: ด,    ASCII: 3604,    Unicode: U+0E14,    Bytes: 3
// Index: 22,   Rune: ี,     ASCII: 3637,    Unicode: U+0E35,    Bytes: 3
```

**Package ที่เกี่ยวกับ String**

1. `strings` – ใช้จัดการกับข้อความ

   ```go
   package main

   import (
    "fmt"
    "strings"
   )

   func main() {
    // ตรวจสอบว่ามีคำนี้ในข้อความหรือไม่ case sensitive
    result1 := strings.Contains("Hello Gopher", "go")
    fmt.Println(result1) // false

    // นับคำที่ต้องการหาว่ามีกี่คำในข้อความ
    result2 := strings.Count("สวัสดีชาวโก", "ดี")
    fmt.Println(result2) // 1

    // ตรวจสอบคำขึ้นต้น
    result3 := strings.HasPrefix("สวัสดีชาวโก", "สวั")
    fmt.Println(result3) // true

    // ตรวจสอบคำลงท้าย
    result4 := strings.HasSuffix("สวัสดีชาวโก", "โก")
    fmt.Println(result4) // true

    // ต่อข้อความจาก []string
    result5 := strings.Join([]string{"สวัสดี", "ชาวโก"}, "_")
    fmt.Println(result5) // สวัสดี_ชาวโก

    // แปลงข้อความเป็นตัวพิมพ์ใหญ่ทั้งหมด
    result6 := strings.ToUpper("Hello Gopher")
    fmt.Println(result6) // HELLO GOPHER

    // แปลงข้อความเป็นตัวพิมพ์เล็กทั้งหมด
    result7 := strings.ToLower("Hello Gopher")
    fmt.Println(result7) // hello gopher
   }
   ```

2. `strconv` – แปลงข้อความเป็นตัวเลข และกลับกัน

   ```go
   package main

   import (
    "fmt"
    "strconv"
   )

   func main() {
    // แปลงเป็น float ต้องระบุขนาดเสมอ (32/64)
    f, _ := strconv.ParseFloat("3.1415", 64)
    fmt.Println(f) // 3.14

    // แปลงเป็น int ต้องระบุเลขฐานของข้อความที่จะแปลงด้วย
    i, _ := strconv.ParseInt("-42", 10, 64)
    fmt.Println(i) // -42

    // แปลงเป็น int ต้องระบุเลขฐานของข้อความที่จะแปลงด้วย
    u, _ := strconv.ParseUint("42", 10, 64)
    fmt.Println(u) // 42

    // แปลงเป็น bool
    b, _ := strconv.ParseBool("true")
    fmt.Println(b) // true

    // แปลงข้อความเป็นจัวเลข
    i2, _ := strconv.Atoi("-42")
    fmt.Println(i2) // -42

    // แปลงตัวเลขเป็นข้อความ
    s := strconv.Itoa(-42)
    fmt.Println(s) // "-42"
   }
   ```

---

### ตัวชี้หน่วยความจำ (Pointer)

Pointer คือ ตัวแปรที่เก็บที่อยู่ของหน่วยความจำ (memory address) ของตัวแปรอื่น แทนที่จะเก็บค่าตรงๆ ทำให้สามารถอ้างอิงและเปลี่ยนแปลงค่าที่อยู่ในหน่วยความจำได้โดยตรง

Zero-value คือ `nil`

**<ins>วิธีการใช้งาน Pointer</ins>**

- การประกาศ Pointer ใช้ `*T` (`*`หน้าชนิดของข้อมูล)

  ```go
  var p *int  // ประกาศ pointer ที่ชี้ไปยังค่า int (ค่าเริ่มต้นเป็น nil)
  ```

- การกำหนดค่าให้ Pointer ใช้ `&` หน้าชื่อตัวแปรที่ต้องการเอาค่า memory address ออกมา

  ```go
  x := 10
  p = &x  // ใช้ & เพื่อดึง memory address ของ x
  ```

- การเข้าถึงค่าผ่าน Pointer ใช้ `*` หน้าชื่อตัวแปร Pointer เพื่อเข้าถึงค่า หรือแก้ไขค่าใน address นั้นๆ

  ```go
  fmt.Println(*p)  // ใช้ * เพื่อเข้าถึงค่าของตัวแปรที่ pointer ชี้อยู่ (ผลลัพธ์คือ 10)
  ```

**ตัวอย่างการใช้งาน Pointer:**

```go
package main

import (
 "fmt"
)

func main() {
 name := "Go"

 var s *string // ประกาศ pointer ที่ชี้ไปยังค่า string

 s = &name // เก็บที่อยู่ของตัวแปร name โดยใช้ & ดึงค่า address ออกมา

 fmt.Println("ค่าของ name:", name)      // ผลลัพท์: Go
 fmt.Println("ที่อยู่ของ s:", s)           // ตัวอย่างผลลัพท์: 0xc000012090
 fmt.Println("ค่าที่ pointer s ชี้อยู่:", *s) // ผลลัพท์: Go

 *s = "Golang" // แก้ไขค่าใน address ที่อ้างอิงถึง

 fmt.Println("ค่าของ name หลังจากเปลี่ยนผ่าน pointer:", name) // ผลลัพท์: Golang
}
```

**<ins>Pass by Value vs Pass by Reference</ins>**

ในภาษา Go การส่งค่าผ่านฟังก์ชันจะใช้ **Pass by Value** เป็นค่าเริ่มต้นเสมอ ซึ่งหมายความว่า เมื่อส่งค่าตัวแปรไปให้ฟังก์ชัน จะเป็นการส่งค่าที่ถูกทำสำเนา (copy) ไป ทำให้ฟังก์ชันสามารถทำงานกับสำเนาของค่าตัวแปรนั้นได้ โดยที่ตัวแปรต้นฉบับจะไม่ถูกแก้ไข

แต่หากต้องการให้ฟังก์ชันสามารถแก้ไขค่าของตัวแปรต้นฉบับได้ ต้องส่งค่าแบบ **Pass by Reference** หรือส่งเป็น **pointer** (ตัวชี้หน่วยความจำ) แทน

- **Pass by Value**

  เมื่อใช้ **Pass by Value** การส่งค่าจะทำการคัดลอกค่าไปยังฟังก์ชันใหม่ ดังนั้นการเปลี่ยนแปลงค่าภายในฟังก์ชันจะไม่มีผลต่อตัวแปรต้นฉบับ

  ```go
  package main

  import "fmt"

  func changeValue(x int) {
      x = 100  // เปลี่ยนค่า x ภายในฟังก์ชัน
      fmt.Println("ค่าภายในฟังก์ชัน:", x)
  }

  func main() {
      a := 10
      changeValue(a)
      fmt.Println("ค่าภายนอกฟังก์ชัน:", a)  // ค่า a ไม่เปลี่ยนแปลง
  }

  // ผลลัพธ์:
  // ค่าภายในฟังก์ชัน: 100
  // ค่าภายนอกฟังก์ชัน: 10
  ```

- **Pass by Reference**

  เมื่อใช้ **Pass by Reference** การส่งค่าจะส่ง **pointer** ของตัวแปรไป ซึ่งหมายความว่า ฟังก์ชันจะสามารถแก้ไขค่าของตัวแปรต้นฉบับได้โดยตรง

  ```go
  package main

  import "fmt"

  func changeValue(x *int) {
      *x = 100  // ใช้ pointer เพื่อเปลี่ยนค่าของตัวแปรต้นฉบับ
      fmt.Println("ค่าภายในฟังก์ชัน:", *x)
  }

  func main() {
      a := 10
      changeValue(&a)  // ส่ง pointer ของ a ไป
      fmt.Println("ค่าภายนอกฟังก์ชัน:", a)  // ค่า a ถูกเปลี่ยนแปลง
  }

  // ผลลัพธ์:
  // ค่าภายในฟังก์ชัน: 100
  // ค่าภายนอกฟังก์ชัน: 100
  ```

---

### ชนิดข้อมูลเชิงโครงสร้าง (Composite Types)

**Composite Types** หรือ **ประเภทข้อมูลเชิงโครงสร้าง** ในภาษา Go คือประเภทข้อมูลที่สามารถรวมหลายๆ ค่า (หรือหลายๆ ฟิลด์) ไว้ในตัวเดียว ซึ่งช่วยให้สามารถเก็บข้อมูลที่มีหลายลักษณะ เช่น การเก็บข้อมูลที่ประกอบไปด้วยค่าหลายประเภทในตัวเดียว เช่น ชื่อ, อายุ, ที่อยู่ เป็นต้น

ประเภทข้อมูลที่ถือเป็น **Composite Types** ในภาษา Go ได้แก่:

- Array
- Slice
- Struct
- Map

---

### Array

**Array** คือประเภทข้อมูลที่เก็บชุดของค่า (หรือสมาชิก) ที่มีขนาดคงที่ (fixed size) และประเภทเดียวกัน

Zero-value คือ เป็นไปตามประเภทข้อมูลนั้น ๆ

```go
package main

import "fmt"

func main() {
 var nums [4]int = [4]int{1, 2, 3} // กำหนด array ขนาด 4

 fmt.Println("Array:", nums) // ผลลัพท์: Array: [1 2 3 0] (0 คือ zero value ของ int)
}
```

การเข้าถึงและแก้ไขค่า

```go
fmt.Println(nums[0]) // ผลลัพท์: 1
fmt.Println(nums[1]) // ผลลัพท์: 2
fmt.Println(nums[2]) // ผลลัพท์: 3
fmt.Println(nums[3]) // ผลลัพท์: 0

nums[3] = 4                 // แก้ไขค่าในตำแหน่งที่ 3
fmt.Println("Array:", nums) // ผลลัพท์: Array: [1 2 3 4]
```

การวนลูปด้วย `for range`

```go
for i, v := range nums {
    fmt.Println(i, v) // i คือ index, v คือ value
}
```

---

### Slice

**Slice** คือประเภทข้อมูลที่คล้ายกับอาเรย์ แต่มีขนาดที่ยืดหยุ่น สามารถเพิ่มหรือลดขนาดได้ตามต้องการ และไม่ต้องกำหนดขนาดล่วงหน้า

Zero-value คือ `nil`

```go
var nums []int // จะมีค่าเป็น zero value คือ nil ยังใช้งานไม่ได้

// ต้อง allocate memory ให้ก่อนถึงจะใช้งานได้
nums = make([]int, 4) // จะต้องระบุขนาดเริ่มต้นให้ก่อน ทุกตำแหน่งจะได้ค่า zero value ของ type นั้นๆ
// []int{0, 0, 0, 0}

// หรือสร้างแบบว่างๆ
nums := []int{}

// หรือ จะใส่ค่าไปตั้งแต่ประกาศเลยก็ได้
nums := []int{1, 2, 3}
```

การเพิ่มค่าลงใน slice ใช้ `append()`

```go
package main

import (
 "fmt"
)

func main() {
 nums := []int{1, 2, 3}

 nums = append(nums, 4)

 fmt.Printf("%#v\n", nums)

}

// ผลลัพธ์:
// []int{1, 2, 3, 4}
```

การดูขนาด ใช้ `len()` และดูพื้นที่สำหรับเก็บข้อมูล ใช้ `cap()`

```go
package main

import (
 "fmt"
)

func main() {
 nums := []int{1, 2, 3}

 nums = append(nums, 4)

 y := len(nums)
 z := cap(nums)

 fmt.Printf("%#v, len=%v, cap=%v", nums, y, z)

}

// ผลลัพธ์:
// []int{1, 2, 3, 4}, len=4, cap=6
```

การ slice ข้อมูล ใช้ `[index_เริ่มต้น:index_ที่ต้องการ+1]`

```go
//          0   1   2   3   4   5   6   7   8
x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

// เอาตั้งแต่ 0 จนถึงตัวสุดท้าย
y := x[0:]

// หรือเขียนแบบนี้ก็ได้
y := x[:]

// ถ้าต้องการ 30, 40, 50
y :=x[2:5] // ค่า 30 คือ index ที่่ 2 ส่วน 50 คือ index ที่ 4 ดังนั้นต้อง + 1 = 5

// ถ้าต้องการต้องแต่เริ่มต้น จนถึงตำแหน่งที่ต้องการ สามารถละ index เริ่มต้นได้
y :=x[:5]
```

> 💡 เปลี่ยน array เป็น slice ง่ายๆ ด้วย `slice := array[:]` แต่เป็นการแชร์ array กัน ถ้าแก้ไขที่ตำแหน่งเดียวกันจะเปลี่ยนทั้งคู่ เพราะ slice จะเป็น memory address ของ array

การลบข้อมูลออกจาก slice ใช้วิธี slice ข้อมูลออกแบบ 2 ก้อน แล้วนำมาต่อกันใหม่ เช่น

```go
words := []string{"A", "B", "C", "D", "E"}

// ถ้าต้องการจะลบ "C" ออกไป ซึ่งก็คือตำแหน่งที่ 2 -> words[:2]
// จะต้องได้ {"A", "B"} + {"D", "E"}
// แต่ไม่สามารถใส่ words[3:] ลงไปตรงๆ ได้
// จะต้องใช้ spread operator แทนแบบนี้
words = append(words[:2], words[3:]...)

fmt.Println(words) // ผลลัพธ์: [A B D E]
```

การวนลูปด้วย `for range`

```go
nums := []int{1, 2, 3}
for i, v := range nums {
    fmt.Println(i, v) // i คือ index, v คือ value
}
```

---

### Map

Map คือประเภทข้อมูลที่เก็บคู่ของคีย์และค่า (key-value pairs) ซึ่งสามารถใช้คีย์ในการค้นหาค่าที่เกี่ยวข้องกับคีย์นั้นๆ ได้

Zero-value คือ `nil`

```go
// ถ้าประกาศขึ้นมาลอยๆ จะมีค่าเป็น zero value คือ nil
var m map[string]string

m := make(map[string]string)
// หรือ
m := map[string]string{}

// หรือ จะใส่ค่าไปตั้งแต่ประกาศเลยก็ได้
m := map[string]string{
 "a": "apple",
  "b": "banana", // ต้องปิดด้วย , เสมอ
}

// เข้าถึงค่าโดยใช้ key
fmt.Println(m["a"]) // apple
```

⚠️ การเข้าถึงค่าโดยใช้ key ถ้าที่ไม่มีอยู่จริงจะได้ zero value กลับมา ทำให้เกิดปัญหาว่า key นั้นมีจริง และมีค่าตามนั้น หรือ key นั่นไม่มี ดังนั้นจะต้องตรวจสอบก่อน

```go
fmt.Printf("%v\n", m["c"])
// ""

v, ok := m["c"]
if ok {
  // แสดงค่าของ key "c"
}
```

การแก้ไขข้อมูล

```go
m["b"] = "berry"

fmt.Printf("%#v\n", m)

// map[string]string{"a":"apple", "b":"berry"}
```

การเพิ่มข้อมูล

```go
m["c"] = "cranberry"

fmt.Printf("%#v\n")

// map[string]string{"a":"apple", "b":"banana", "c":"cranberry"}
```

การลบข้อมูล

```go
delete(m, "b")

fmt.Printf("%#v\n", m)

// map[string]string{"a":"apple", "c":"cranberry"}
```

การวนลูปด้วย `for range`

```go
m := map[string]string{
 "a": "apple",
  "b": "banana",
}

for k, v := range m {
 fmt.Println(k, v) // จะได้ key แทน index
}
```

---

### Struct

Struct คือประเภทข้อมูลที่เก็บค่าหลายประเภท โดยค่าที่เก็บใน struct อาจจะเป็นข้อมูลที่มีประเภทต่างกัน เช่น ตัวเลข, สตริง หรือประเภทอื่นๆ

```go
package main

import "fmt"

type Rectangle struct {
 Width  float64
 Height float64
}

func main() {
 rec := Rectangle{} // จะมีค่าเป็น empty -> {}

 // หรือจะกำหนดค่าไปเลยก็ได้

 rec = Rectangle{
  Width:  10,
  Height: 20,
 }

 // การอ่านค่า
 fmt.Println(rec.Width) // 10

 // การกำหนด/แก้ไขค่า
 rec.Width = 30
 fmt.Println(rec.Width) // 30
}
```

**<ins>Method</ins>**

ในภาษา Go **Method** คือฟังก์ชันที่ถูกผูกไว้กับ `struct` หรือประเภทข้อมูลที่กำหนด โดย method จะมีการกำหนด **receiver** (ตัวรับค่า) ซึ่งเป็นตัวแปรที่ใช้แทน `struct` ที่ method นั้นจะทำงานด้วย

ก่อนที่เราจะเข้าใจเรื่อง method ของ `struct` ให้เรามาดูฟังก์ชันธรรมดากันก่อน ฟังก์ชันธรรมดาคือฟังก์ชันที่ไม่มี receiver ซึ่งหมายความว่า ฟังก์ชันเหล่านี้ไม่ได้ผูกกับตัวแปรใดๆ โดยตรง

```go
package main

import "fmt"

type Rectangle struct {
 Width  float64
 Height float64
}

// normal fucntion
func Area(rec Rectangle) float64 {
 return rec.Width * rec.Height
}

func main() {
 rec := Rectangle{
  Width:  10,
  Height: 20,
 }

 // fucntion style
 fmt.Println(Area(rec))
}
```

การกำหนด method ให้กับ `struct` หรือประเภทข้อมูลอื่น ๆ จะใช้รูปแบบดังนี้:

```go
func (receiver ReceiverType) MethodName() ReturnType {
    // การทำงานของ method
}
```

**ตัวอย่างการใช้ Method:**

```go
package main

import "fmt"

type Rectangle struct {
 Width  float64
 Height float64
}

// recevier function
func (rec Rectangle) Area() float64 {
 return rec.Width * rec.Height
}

func main() {
 rec := Rectangle{
  Width:  10,
  Height: 20,
 }

 // method style
 fmt.Println(rec.Area())
}
```

**<ins>Receiver Type</ins>**

Receiver Type สามารถใช้ได้ทั้งในรูปแบบของ **Value Receiver** หรือ **Pointer Receiver** ซึ่งการเลือกใช้ขึ้นอยู่กับความต้องการของการจัดการกับข้อมูลใน `struct`

- **Value Receiver:**

  เมื่อใช้ value receiver, Go จะทำการคัดลอกค่าของ `struct` นั้นๆ มาใช้ใน method ซึ่งหมายความว่า method จะไม่สามารถแก้ไขค่าภายใน `struct` ได้

  ตัวอย่าง:

  ```go
  package main

  import "fmt"

  type Rectangle struct {
   Width  float64
   Height float64
  }

  // Value receiver
  func (r Rectangle) setWidth(w float64) {
   r.Width = w // การเปลี่ยนแปลงนี้จะไม่ส่งผลกับตัวแปร rect ที่ถูกเรียก
  }

  func main() {
   rect := Rectangle{Width: 10, Height: 20}
   rect.setWidth(20)
   fmt.Println("Width after set:", rect.Width) // ผลลัพธ์: 10
  }
  ```

- **Pointer Receiver:**

  เมื่อใช้ pointer receiver, method จะรับพอยน์เตอร์ของ `struct` ซึ่งทำให้ method สามารถแก้ไขข้อมูลภายใน `struct` ได้

  ตัวอย่าง:

  ```go
  package main

  import "fmt"

  type Rectangle struct {
   Width  float64
   Height float64
  }

  // Pointer receiver
  func (r *Rectangle) setWidth(w float64) {
      r.Width = w // การเปลี่ยนแปลงนี้จะส่งผลกับตัวแปร rect ที่ถูกเรียก
  }

  func main() {
      rect := Rectangle{Width: 10, Height: 5}
      rect.setWidth(20)
      fmt.Println("Width after set:", rect.Width) // ผลลัพธ์: 20
  }
  ```

การเลือกใช้ Value Receiver หรือ Pointer Receiver

- **Value Receiver:** ใช้เมื่อคุณไม่ต้องการให้ method แก้ไขข้อมูลของ `struct` หรือเมื่อ `struct` เป็นค่าที่เล็ก (เช่น `int`, `string`, หรือ `float`)
- **Pointer Receiver:** ใช้เมื่อคุณต้องการให้ method แก้ไขข้อมูลใน `struct` หรือเมื่อ `struct` เป็นค่าที่ใหญ่

**<ins>Type Embedding</ins>**

ใน Go, **Type Embedding** หรือ **Composition** คือแนวทางการออกแบบที่ช่วยให้ `struct` สามารถฝัง (embed) `struct` อื่นได้ ซึ่งทำให้ `struct` ใหม่สามารถใช้ฟังก์ชันและพฤติกรรมจาก `struct` ที่ถูกฝังโดยตรง โดยไม่จำเป็นต้องประกาศฟังก์ชันใหม่

การใช้ **Type Embedding** ช่วยให้ Go มีลักษณะคล้ายกับ **Inheritance** (การสืบทอด) ในภาษาอื่น ๆ เช่น C++ หรือ Java แต่ใน Go, การใช้ embedding จะเป็นการ **composition** ซึ่งเน้นความยืดหยุ่นและการบำรุงรักษาที่ง่าย

```go
package main

import "fmt"

type Shape struct {
 Name string
}

type Rectangle struct {
 Width  float64
 Height float64
 Shape  // Embedding Shape
}

func main() {
 rec := Rectangle{
  Width:  10,
  Height: 20,
  Shape:  Shape{Name: "Rectangle"},
 }

 fmt.Println(rec.Shape.Name) // ผลลัพธ์: Rectangle
 // คุณสมบัติการ promoted fields
 fmt.Println(rec.Name)       // ผลลัพธ์: Rectangle
}
```

> 💡 หากมี method ชื่อเดียวกันใน `struct` method จาก `struct` ที่ประกาศล่าสุดจะถูกใช้

---

## 4. ออกแบบโค้ดให้ยืดหยุ่นด้วย Interface

การใช้ **Interface** ในภาษา Go เป็นเครื่องมือที่ช่วยให้โค้ดของเรายืดหยุ่นและสามารถทำงานได้กับหลายประเภทข้อมูล โดยไม่ต้องระบุชนิดข้อมูลเฉพาะล่วงหน้า **interface** คือชนิดข้อมูลที่สามารถเก็บค่าได้หลายประเภท เช่น `int`, `string`, `struct` หรือแม้กระทั่งฟังก์ชัน และเราสามารถใช้มันเพื่อกำหนดพฤติกรรมที่ต้องการจากประเภทข้อมูลต่างๆ

- [Empty Interface (`interface{}`)](#empty-interface-interface)
- [Interface ที่มี method](#interface-ที่มี-method)
- [Type Assertion (การยืนยันประเภทข้อมูล)](#type-assertion-การยืนยันประเภทข้อมูล)
- [Nil Interface (กรณีที่ interface มีค่าเป็น nil)](#nil-interface-กรณีที่-interface-มีค่าเป็น-nil)

---

### Empty Interface (`interface{}`)

`interface{}` คือ **empty interface** ที่สามารถเก็บค่าได้ทุกรูปแบบ โดยไม่ต้องระบุชนิดข้อมูลที่ชัดเจน

**Zero-value:** `nil`

**ตัวอย่าง:**

```go
package main

import "fmt"

// Use Case 1: เมื่อต้องการรับค่าใดๆ ที่ไม่รู้ประเภทล่วงหน้า
// ฟังก์ชันที่ใช้รับค่าเป็น Empty Interface (interface{})
func printAnything(value interface{}) {
    fmt.Println(value)
}

func main() {
    printAnything("Hello, World!")   // รับค่า string
    printAnything(123)               // รับค่า int
    printAnything(true)              // รับค่า bool

    // Use Case 2: ต้องการรับค่าหลายประเภทไว้ในตัวแปรเดียว
    var x interface{}  // สามารถเก็บข้อมูลได้ทุกชนิด
  x = 10             // เก็บ int
  x = "Hello"        // เก็บ string
}

// **ผลลัพธ์:
//** Hello, World!
**//** 123
**//** true
```

---

### Interface ที่มี method

ใช้เพื่อกำหนดรูปแบบการทำงานของ Struct ที่ Implement Interface นั้นๆ

ใน Go จะเป็น Implicit Implementation คือ ถ้า Struct มีเมธอดที่ตรงกับ Interface ทุกตัว ถือว่า Implement Interface นั้นทันที

**ตัวอย่าง:**

```go
package main

import "fmt"

// สร้าง interface ที่กำหนด method Speak
type Speaker interface {
    Speak() string
}

// สร้าง struct ที่ implement interface Speaker
type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

// ฟังก์ชันที่รับ argument เป็น interface Speaker
func introduce(speaker Speaker) {
    fmt.Println(speaker.Speak())
}

func main() {
    person := Person{Name: "Alice"}
    introduce(person) // จะใช้ method Speak ที่อยู่ใน struct Person
}

// **ผลลัพธ์:
//** Hello, my name is Alice
```

**คำอธิบาย:**

- `Speaker` เป็น interface ที่มี method `Speak()` ซึ่งทุกๆ ชนิดที่ implement interface นี้จะต้องมี method `Speak()` ที่ทำงานบางอย่าง
- `Person` struct มี method `Speak()` ที่ทำให้ `Person` เป็นไปตามข้อกำหนดของ `Speaker` interface
- ฟังก์ชัน `introduce` รับ parameter ที่เป็น `Speaker` ซึ่งสามารถรับค่าทุกประเภทที่มี method `Speak`

---

### Type Assertion (การยืนยันประเภทข้อมูล)

Type assertion ใช้เพื่อ **ตรวจสอบหรือแปลง** ข้อมูลที่เก็บใน `interface{}` เป็นประเภทข้อมูลที่เราต้องการ

**รูปแบบการใช้งาน:**

```go
value, ok := iface.(Type)
```

**หลักการทำงาน:**

- แปลงค่าภายใน `iface` เป็นประเภทที่ต้องการ โดยใช้เครื่องหมาย `.(Type)`
- ถ้าไม่สามารถแปลงได้จะได้รับค่าผลลัพธ์เป็น `false` ในตัวแปร `ok`

**ตัวอย่าง:**

```go
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
    checkType("Hello, Go!")  // String
    checkType(42)            // Integer
    checkType(3.14)          // Unknown type
}

// ผลลัพธ์:
// This is a string: Hello, Go!
// This is an integer: 42
// Unknown type
```

**คำอธิบาย:**

- ทำการ **type assertion** เพื่อเช็คประเภทของค่าที่รับเข้ามาใน `checkType`
- ใช้ `value.(string)` เพื่อดูว่าค่าที่รับมาเป็นประเภท `string` หรือไม่
- ถ้าไม่ใช่ `string` เราก็เช็คต่อว่าเป็น `int` หรือไม่
- ถ้าไม่ตรงกับประเภทใดๆ ที่เราตรวจสอบ จะเข้าสู่สภาวะ "Unknown type"

---

### Nil Interface (กรณีที่ interface มีค่าเป็น nil)

Interface สามารถมีค่าเป็น `nil` ได้ แต่จะต้องระมัดระวังเมื่อใช้ เพราะ `interface` ที่เป็น `nil` จะมีค่าเป็น `nil` ทั้งประเภทและค่า

ถ้าหาก `interface` ถูกประกาศแต่ไม่ได้เก็บค่าใดๆ (เป็น `nil`) เราจะไม่สามารถเรียกใช้ method หรือ field ได้

**ตัวอย่าง:**

```go
package main

import "fmt"

func checkNil(value interface{}) {
    if value == nil {
        fmt.Println("The interface is nil")
    } else {
        fmt.Println("The interface is not nil")
    }
}

func main() {
    var i interface{}
    checkNil(i)  // ค่าเป็น nil

    i = 10
    checkNil(i)  // ค่าไม่เป็น nil
}

// ผลลัพธ์:
// The interface is nil
// The interface is not nil
```

---

## 5. ทำความรู้จักกับ Generic ในภาษา Go

Generics เป็นความสามารถที่เพิ่มเข้ามาใน Go ตั้งแต่เวอร์ชัน 1.18 ช่วยให้เราสามารถสร้างฟังก์ชัน, โครงสร้างข้อมูล (struct), และอินเทอร์เฟซที่ทำงานกับหลายชนิดข้อมูล (types) ได้โดยไม่ต้องเขียนโค้ดซ้ำ

- [ก่อนหน้าที่จะมี Generics](#ก่อนหน้าที่จะมี-generics)
- [การใช้ Generic Type กับฟังก์ชัน](#การใช้-generic-type-กับฟังก์ชัน)
- [การใช้ Generic Type กับ Struct](#การใช้-generic-type-กับ-struct)
- [การใช้ Type Constraint ใน Generics](#การใช้-type-constraint-ใน-generics)
- [ข้อดีของการใช้ Generics](#ข้อดีของการใช้-generics)

---

### ก่อนหน้าที่จะมี Generics

ก่อน Go 1.18 การสร้างฟังก์ชันที่รองรับหลายชนิดข้อมูลต้องใช้ `interface{}` และการทำ **Type Assertion** ซึ่งทำให้โค้ดดูยุ่งยากและอาจเกิดข้อผิดพลาดในช่วง **Runtime** เนื่องจากไม่มีการตรวจสอบประเภทข้อมูล (type-safety) ในระหว่างการคอมไพล์

**ตัวอย่าง:**

```go
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
    fmt.Println(Sum([]interface{}{1, 2, 3}))        // 6
    fmt.Println(Sum([]interface{}{1.5, 2.5}))      // 4.0
    fmt.Println(Sum([]interface{}{"hello", 2, 3})) // error
}
```

ปัญหาของการใช้ `interface{}`

1. **ไม่มี Type-safety**: ต้องแปลงประเภทเองเมื่อใช้งาน
2. **ประสิทธิภาพต่ำกว่า**: มี Overhead จากการแปลงประเภท

---

### การใช้ Generic Type กับฟังก์ชัน

เราสามารถสร้างฟังก์ชันที่รองรับหลายประเภทข้อมูลโดยใช้ **Type Parameter** ซึ่งจะกำหนดชื่อประเภทข้อมูลในวงเล็บเหลี่ยม `[]` เพื่อให้ฟังก์ชันสามารถทำงานกับประเภทต่างๆ ได้

**ตัวอย่าง:**

```go
package main

import "fmt"

// T คือ Type Parameter ที่รองรับ int และ float64
func Sum[T int | float64](nums []T) T {
    var total T
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(Sum([]int{1, 2, 3}))       // 6
    fmt.Println(Sum([]float64{1.5, 2.5})) // 4.0
}
```

**คำอธิบาย:**

- `T` เป็นตัวแทนของประเภท (type parameter)
- `T int | float64` กำหนดข้อจำกัด (constraint) ว่า `T` ต้องเป็น `int` หรือ `float64`

---

### การใช้ Generic Type กับ Struct

เราสามารถสร้างโครงสร้างข้อมูล (`struct`) ที่รองรับหลายประเภทข้อมูลได้ โดยการใช้ **Generic Type** ในการกำหนดประเภทข้อมูลของฟิลด์ภายใน `struct`

**ตัวอย่าง:**

```go
package main

import "fmt"

// โครงสร้างข้อมูลแบบ Generic
type Box[T any] struct {
    value T
}

// เมธอดที่ใช้ Generic Type
func (b Box[T]) GetValue() T {
    return b.value
}

func main() {
    intBox := Box[int]{value: 100}
    stringBox := Box[string]{value: "Hello Go"}

    fmt.Println(intBox.GetValue())   // 100
    fmt.Println(stringBox.GetValue()) // Hello Go
}
```

**คำอธิบาย:**

- `Box[T any]` คือ โครงสร้างที่รองรับทุกประเภท (`any` คือ type constraint ที่รองรับทุกชนิดข้อมูล)
- `(b Box[T]) GetValue() T` คือ เมธอดที่ใช้ Generic Type

---

### การใช้ Type Constraint ใน Generics

การใช้ **Type Constraint** ใน Generics ช่วยให้เราสามารถจำกัดประเภทข้อมูล (`T`) ที่ใช้ในฟังก์ชันหรือ `struct` ให้เป็นไปตามที่เราต้องการ โดยสามารถกำหนดได้ว่า `T` ต้องเป็นประเภทใดประเภทหนึ่ง หรือประเภทที่รองรับพฤติกรรมบางอย่าง เช่น การใช้ method เฉพาะหรือการทำงานร่วมกับชนิดข้อมูลที่สอดคล้องกันใน Go

ตัวอย่าง:

```go
package main

import (
    "fmt"
)

// Number คือ Type Constraint ที่กำหนดให้ T ต้องเป็น int หรือ float64 เท่านั้น
type Number interface {
    int | float64
}

// ใช้ Generics พร้อม Type Constraint
func Sum[T Number](nums []T) T {
    var total T
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(Sum([]int{1, 2, 3}))        // 6
    fmt.Println(Sum([]float64{1.5, 2.5}))  // 4.0
    // fmt.Println(Sum([]string{"a", "b"})) ❌ Error: string ไม่ได้อยู่ใน Number constraint
}
```

**คำอธิบาย**:

- `Number` เป็น `interface` ที่กำหนดข้อจำกัดให้รองรับ `int`, `float64`
- ฟังก์ชัน `Sum` ใช้ `T Number` เพื่อจำกัดเฉพาะประเภทข้อมูลที่รับเข้ามาได้

---

### ข้อดีของการใช้ Generics

1. **ลดการทำซ้ำโค้ด (Code Reusability)**: สามารถสร้างฟังก์ชันและโครงสร้างข้อมูลที่ทำงานได้กับหลายประเภท
2. **เพิ่มความปลอดภัย (Type Safety)**: ประเภทข้อมูลถูกตรวจสอบในขณะคอมไพล์
3. **ประสิทธิภาพที่ดีขึ้น**: ลดการใช้ `interface{}` ที่มี Overhead สูง
4. **ใช้กับ Struct, Method และ Interface ได้**

---

## 6. การทำงานแบบ Concurrency

**Concurrency** ใน Go ช่วยให้สามารถประมวลผลหลายๆ อย่างพร้อมกันได้ โดยไม่ต้องรอให้กระบวนการหนึ่งเสร็จสิ้นก่อนถึงจะเริ่มกระบวนการถัดไป ซึ่งช่วยเพิ่มประสิทธิภาพและการทำงานที่เร็วขึ้น

เครื่องมือสำคัญที่ใช้ในการทำงานแบบ Concurrency ได้แก่:

- [Goroutine](#goroutine): ใช้สำหรับรันฟังก์ชันหรือกระบวนการแบบขนาน (parallel) ที่ทำงานพร้อมกันหลายตัว
- [WaitGroup](#waitgroup): ใช้เพื่อรอให้ Goroutine ทุกตัวที่เริ่มทำงานเสร็จสิ้นก่อนที่จะดำเนินการขั้นตอนถัดไป
- [Channel](#channel): ใช้สำหรับการส่งข้อมูลและสื่อสารระหว่าง Goroutine
- [Mutex](#mutex) ใช้ในการควบคุมการเข้าถึงทรัพยากรที่ใช้ร่วมกัน เพื่อป้องกันปัญหา **Race Condition** ที่เกิดจากการเข้าถึงข้อมูลพร้อมกัน
- [Select Statement](#select-statement): ใช้ในการจัดการหลายๆ Channel พร้อมกัน เพื่อรอรับข้อมูลจาก Channel ใดที่พร้อมใช้งาน

---

### Goroutine

**Goroutine** คือ **lightweight thread** ใน Go ที่ช่วยให้เราสามารถรันโค้ดแบบขนาน (**concurrent execution**) ได้อย่างมีประสิทธิภาพ

✅ เบากว่า OS Thread ช่วยประหยัดทรัพยากร และทำให้โปรแกรมทำงานได้เร็วขึ้น

✅ มีการจัดการโดย Go Runtime

✅ ทำให้แอปพลิเคชันรองรับ **Concurrency** ได้ง่าย

การเริ่มใช้งาน Goroutine สามารถทำได้ง่ายๆ โดยการใช้คำสั่ง `go` ก่อนชื่อฟังก์ชัน เช่น:

```go
go myFunction()
```

ฟังก์ชันนี้จะถูกเรียกใน Goroutine ใหม่ ซึ่งจะทำงานพร้อมกับส่วนอื่นของโปรแกรม ทำให้โปรแกรมสามารถทำงานหลายๆ งานพร้อมกันได้อย่างมีประสิทธิภาพ

**ตัวอย่าง:**

```go
package main

import (
 "fmt"
 "time"
)

// ฟังก์ชันที่จะรันใน Goroutine
func worker(id int) {
 fmt.Printf("Worker %d started\n", id)
}

func main() {
 for i := 1; i <= 3; i++ {
  go worker(i) // เรียกใช้ฟังก์ชันใน Goroutine
 }

 fmt.Println("main function running")

 time.Sleep(500 * time.Millisecond)
 fmt.Println("All workers done!")
}
```

**ผลลัพธ์ (อาจเปลี่ยนแปลงได้ในแต่ละครั้ง):**

```
main function running
Worker 2 started
Worker 3 started
Worker 1 started
All workers done!
```

> 📝 หมายเหตุ: ถ้าไม่มี `time.Sleep` Goroutine อาจไม่ทันทำงาน เพราะ `main()` จบก่อน

---

### WaitGroup

ใช้ `sync.WaitGroup` ใช้เพื่อรอให้ Goroutine ทุกตัวที่เริ่มทำงานเสร็จสิ้นก่อนที่จะดำเนินการขั้นตอนถัดไป

**ตัวอย่าง:**

```go
package main

import (
 "fmt"
 "sync"
)

// ฟังก์ชันที่จะรันใน Goroutine
func worker(id int, wg *sync.WaitGroup) {
 defer wg.Done() // บอก WaitGroup ว่า Goroutine เสร็จแล้ว
 fmt.Printf("Worker %d started\n", id)
}

func main() {
 var wg sync.WaitGroup
 for i := 1; i <= 3; i++ {
  wg.Add(1) // เพิ่มจำนวน Goroutine
  go worker(i, &wg)
 }

 fmt.Println("main function running")

 wg.Wait() // รอจน Goroutine ทั้งหมดเสร็จ
 fmt.Println("All workers done!")
}
```

**ผลลัพธ์ (อาจเปลี่ยนแปลงได้ในแต่ละครั้ง):**

```
main function running
Worker 1 started
Worker 2 started
Worker 3 started
All workers done!
```

---

### Channel

**Channel** ใช้สำหรับการส่งข้อมูลและสื่อสารระหว่าง Goroutine

**ตัวอย่าง:**

```go
package main

import "fmt"

func sendData(ch chan string) {
    ch <- "Hello from Goroutine" // ส่งข้อมูลผ่าน Channel
}

func main() {
    ch := make(chan string) // สร้าง Channel
    go sendData(ch)         // เรียก Goroutine
    msg := <-ch             // รอรับข้อมูลจาก Channel
    fmt.Println(msg)
}
```

**ผลลัพธ์:**

```
Hello from Goroutine
```

> 📝 หมายเหตุ:
>
> - `make(chan string)` ใช้สร้าง Channel ที่รับข้อมูลประเภท `string`
> - Channel จะบล็อก (รอ) จนกว่ามีการรับ-ส่งข้อมูล

**Buffered Channel** สามารถเก็บข้อมูลได้หลายค่าพร้อมกัน โดยไม่บล็อกทันที

**ตัวอย่าง:**

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3) // Buffered Channel ที่รองรับ 3 ค่า

    ch <- 1
    ch <- 2
    ch <- 3

    fmt.Println(<-ch) // ดึงค่าออก
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

```

**ผลลัพธ์:**

```
1
2
3
```

> 📝 หมายเหตุ: ถ้าใส่ข้อมูลเกินความจุ Channel จะเกิด Deadlock

---

### Mutex

**Mutex** (Mutual Exclusion) ใช้ในการควบคุมการเข้าถึงทรัพยากรที่ใช้ร่วมกัน เพื่อป้องกันปัญหา **Race Condition** ซึ่งเกิดจากการเข้าถึงและแก้ไขข้อมูลพร้อมกันโดยหลายๆ Goroutine

**Race Condition คืออะไร?**

**Race Condition** เกิดขึ้นเมื่อหลาย Goroutine พยายามเข้าถึงและแก้ไขข้อมูลเดียวกันพร้อมกันโดยไม่ได้มีการควบคุม ซึ่งอาจทำให้ข้อมูลที่ได้ไม่ถูกต้องหรือผิดพลาด

**ตัวอย่าง:**

```go
package main

import (
 "fmt"
 "sync"
)

var (
 count = 0
)

func increment(wg *sync.WaitGroup) {
 defer wg.Done()
 for i := 0; i < 1000; i++ {
  count++ // อัปเดตค่าพร้อมกัน อาจเกิด Race Condition
 }
}

func main() {
 var wg sync.WaitGroup
 for i := 0; i < 10; i++ {
  wg.Add(1)
  go increment(&wg)
 }

 wg.Wait() // รอให้ทุก goroutine เสร็จทำงาน
 fmt.Println("Final count:", count) // ค่าผิดพลาดเพราะ Race Condition
}
```

**ผลลัพธ์ (คาดเดาไม่ได้):**

```
Final count: 7931
```

> 📝 หมายเหตุ: ค่า `count` ควรเป็น `10000` แต่เนื่องจาก Race Condition ทำให้ค่าผิดพลาด

**การแก้ปัญหา Race Condition ด้วย Mutex**

`sync.Mutex` คือเครื่องมือที่ใช้ในการป้องกัน Race Condition โดยการล็อกการเข้าถึงทรัพยากรที่ใช้ร่วมกัน ทำให้ Goroutine อื่นๆ ไม่สามารถเข้าถึงและแก้ไขข้อมูลได้จนกว่าการเข้าถึงที่ล็อกไว้จะเสร็จสิ้น

**ตัวอย่าง:**

```go
package main

import (
 "fmt"
 "sync"
)

var (
 count = 0
 mu    sync.Mutex
)

func increment(wg *sync.WaitGroup) {
 defer wg.Done()
 for i := 0; i < 1000; i++ {
  mu.Lock() // ล็อกก่อนเข้าถึงข้อมูล เพื่อไม่ให้ Goroutine อื่นเข้ามาแก้ไขตัวแปร counter จนกว่าการทำงานจะเสร็จสิ้น
  count++
  mu.Unlock() // ปลดล็อกเมื่อเสร็จเรียบร้อย
 }
}

func main() {
 var wg sync.WaitGroup
 for i := 0; i < 10; i++ {
  wg.Add(1)
  go increment(&wg)
 }

 wg.Wait() // รอให้ทุก goroutine เสร็จทำงาน
 fmt.Println("Final count:", count)
}
```

**ผลลัพธ์:**

```
Final count: 10000
```

---

### Select Statement

การใช้ `select` ช่วยให้โปรแกรมสามารถจัดการกับหลายๆ Channel โดยไม่ต้องรอให้แต่ละ Channel เสร็จสิ้นก่อน จึงทำให้สามารถประมวลผลข้อมูลได้พร้อมกันหลายๆ งาน

**ตัวอย่าง:**

```go
package main

import (
    "fmt"
    "time"
)

func sendData(ch chan string, msg string, delay time.Duration) {
    time.Sleep(delay)
    ch <- msg
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go sendData(ch1, "From Channel 1", 2*time.Second)
    go sendData(ch2, "From Channel 2", 1*time.Second)

    select {
    case msg := <-ch1:
        fmt.Println(msg)
    case msg := <-ch2:
        fmt.Println(msg)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout!")
    }
}
```

**ผลลัพธ์:**

```
From Channel 2
```

---

## 7. การจัดการข้อผิดพลาด (Error Handling)

การจัดการข้อผิดพลาด (Error Handling) เป็นสิ่งสำคัญในภาษา Go เนื่องจาก Go ไม่มีระบบ `try-catch` แต่ใช้แนวทางที่ตรงไปตรงมาในการจัดการข้อผิดพลาด ซึ่งช่วยให้โค้ดมีความชัดเจนและเข้าใจง่าย โดยใน Go เราจะแบ่งการจัดการข้อผิดพลาดออกเป็นหลายส่วน ได้แก่:

- [Error interface](#error-interface): ใช้สำหรับการตรวจสอบข้อผิดพลาดที่เกิดขึ้น
- [Custom Error](#การสร้าง-custom-error-ใน-go): ใช้สำหรับสร้างข้อผิดพลาดที่มีข้อมูลเฉพาะ
- [Panic](#panic): ใช้เมื่อเกิดข้อผิดพลาดที่ไม่สามารถคาดการณ์ได้
- [Recover](#ดักจับ-panic-ด้วย-recover): ใช้เพื่อจับข้อผิดพลาดจาก `panic` และฟื้นฟูการทำงานของโปรแกรม

---

### Error interface

ในภาษา Go ไม่มีการใช้ `try-catch` สำหรับจัดการกับข้อผิดพลาด (error) ซึ่งต่างจากภาษาอื่นๆ แต่ Go ใช้การตรวจสอบข้อผิดพลาด (error handling) อย่างตรงไปตรงมา โดยเราจะต้องจัดการกับข้อผิดพลาดในจุดที่เกิดขึ้นทันที

การจัดการ Error ใน Go ทำได้โดยการตรวจสอบว่ามีค่าผิดพลาด (`error`) หรือไม่ โดยจะต้องตรวจสอบว่า `error` นั้นเป็นค่า `nil` หรือไม่ ซึ่งค่าผิดพลาดใน Go เป็นประเภท `interface` และมีค่าเริ่มต้น (zero value) เป็น `nil` หากไม่มีข้อผิดพลาดเกิดขึ้น

**ขั้นตอนในการจัดการ Error:**

1. **ตรวจสอบค่าผิดพลาด (`error`)** ว่ามีค่าหรือไม่
2. หากมีข้อผิดพลาด ค่า `error` จะต้องมีค่าไม่เท่ากับ `nil` (`!= nil`) ซึ่งแสดงว่ามีข้อผิดพลาดเกิดขึ้น
3. หากไม่มีข้อผิดพลาด (ไม่มีค่า `error` หรือค่าเป็น `nil`) ก็จะสามารถทำงานต่อได้

**ตัวอย่าง:**

```go
package main

import (
 "fmt"
 "strconv"
)

func main() {
 n, err := strconv.Atoi("5s")
 if err != nil {
  // อาจจะ log หรือส่งออกไปให้ที่อื่นจัดการต่อ
  fmt.Println("Error converting string to integer:", err)
  return
 }
 fmt.Println("Number of seconds:", n)
}

// ผลลัพธ์:
// Error converting string to integer: strconv.Atoi: parsing "5s": invalid syntax
```

---

### การสร้าง Custom Error ใน Go

ใน Go เราสามารถสร้าง **Custom Error** เพื่อให้ Error มีข้อมูลที่เป็นประโยชน์มากขึ้น เช่น **Context, Stack Trace หรือ Metadata**

สามารถทำได้หลายวิธี ดังนี้:

1. ใช้ `errors.New()` สำหรับ Error แบบง่าย

   ```go
   package main

   import (
    "errors"
    "fmt"
   )

   func doSomething() error {
    return errors.New("เกิดข้อผิดพลาดบางอย่าง")
   }

   func main() {
    err := doSomething()
    if err != nil {
     fmt.Println("Error:", err)
    }
   }

   // ผลลัพธ์:
   // Error: เกิดข้อผิดพลาดบางอย่าง
   ```

2. ใช้ `fmt.Errorf()` เพื่อเพิ่มข้อความ Context ใน Error

   ```go
   package main

   import (
    "fmt"
   )

   func readFile(filename string) error {
    return fmt.Errorf("ไม่พบไฟล์: %s", filename)
   }

   func main() {
    err := readFile("data.txt")
    if err != nil {
     fmt.Println("Error:", err)
    }
   }

   // ผลลัพธ์:
   // Error: ไม่พบไฟล์: data.txt
   ```

3. ใช้ `type struct` สร้าง Custom Error โดย Implement `Error()` เพื่อกำหนด **Error Code, Message, Context** ได้เอง

   ```go
   package main

   import (
    "fmt"
   )

   // CustomError คือ Struct ที่เก็บข้อมูล Error
   type CustomError struct {
    Code    int
    Message string
   }

   // Implement Method Error() ตาม Interface error
   func (e *CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
   }

   func doSomething() error {
    return &CustomError{Code: 404, Message: "Resource Not Found"}
   }

   func main() {
    err := doSomething()
    if err != nil {
     fmt.Println(err)

     // แปลง Error เป็น CustomError เพื่อดึงค่า Code
     if customErr, ok := err.(*CustomError); ok {
      fmt.Println("Error Code:", customErr.Code)
     }
    }
   }

   // ผลลัพธ์:
   // Error 404: Resource Not Found
   // Error Code: 404
   ```

---

### Panic

`panic` ใช้สำหรับการหยุดการทำงานของโปรแกรมเมื่อเกิดข้อผิดพลาดที่ไม่สามารถแก้ไขได้ โดยทั่วไปจะใช้ `panic` เมื่อเจอข้อผิดพลาดที่ไม่สามารถคาดการณ์ล่วงหน้าได้

การใช้ `panic` จะทำให้โปรแกรมหยุดทำงานทันทีและคืนค่า stack trace ให้เราเห็น

```go
package main

import (
 "fmt"
)

func main() {
 fmt.Println(1)
 fmt.Println(2)
 panic("Fail")
 fmt.Println(3) // ไม่ได้ถูกเรียกใช้
 fmt.Println(4) // ไม่ได้ถูกเรียกใช้
 fmt.Println(5) // ไม่ได้ถูกเรียกใช้
}

// ผลลัพธ์:
// 1
// 2
// panic: Fail
```

---

### ดักจับ Panic ด้วย Recover

เมื่อมีการเรียกฟังก์ชันซ่อนๆ กัน แล้วเกิด panic ขึ้น ระบบจะดูว่าในฟังก์ชันที่เรียกฟังก์ชันที่เกิด panic นั้น ได้มีการจัดการ panic หรือไม่ ถ้าไม่มีก็จะส่ง panic ต่อขึ้นไปเรื่อยๆ จนถึง `main()` ถ้าไม่การการจัดการก็จะจบการทำงานของโปรแกรมไป

ซึ่งเราสามารถดักจับ panic ได้ด้วยการใช้ `recover()`

ตัวอย่าง:

```go
package main

import (
 "fmt"
)

func a() {
 b()
}

func b() {
 panic("Panic in b")
}

func main() {
 a()
 fmt.Println("Completed")
}

// ผลลัพธ์:
// panic: Panic in b
```

แบบนี้จะโปรแกรมจะหยุดการทำงานเพราะเกิด panic แต่ถ้าเพิ่ม `recover()` เข้าไปใน `a()` โปรแกรมก็จะทำงานได้ต่อจนจบ ดังนี้:

```go
package main

import (
 "fmt"
)

func a() {
 defer func() {
  if r := recover(); r != nil {
   fmt.Println("Recover in a:", r)
  }
 }()
 b()
}

func b() {
 panic("Panic in b")
}

func main() {
 a()
 fmt.Println("Completed")
}

// ผลลัพธ์:
// Recover in a: Panic in b
// Completed
```

---

## 8. Unit Testing

Unit Testing เป็นส่วนสำคัญของการพัฒนาซอฟต์แวร์ เพื่อให้แน่ใจว่าฟังก์ชันต่าง ๆ ทำงานถูกต้องตามที่คาดหวัง

ใน Go เราสามารถเขียน Unit Test ได้โดยใช้แพ็กเกจ `testing` ซึ่งเป็นส่วนหนึ่งของ Go standard library

- [การเขียน Unit test](#การเขียน-unit-test)
- [การรัน Unit test](#การรัน-unit-test)
- [การทำ Table-Driven Test](#การทำ-table-driven-test)

---

### การเขียน Unit test

โดยทั่วไป การเขียน Unit Test ใน Go จะมีรูปแบบดังนี้:

- ถ้าไฟล์หลักชื่อ `main.go`
- ไฟล์ทดสอบต้องชื่อ `main_test.go` (ชื่อไฟล์ต้องลงท้ายด้วย `_test.go`)
- ฟังก์ชันของ unit test ส่วนใหญ่จะใช้ prefix ว่า `Test` จะต้องรับพารามิเตอร์ 1 ตัว คือ `(t *testing.T)`
- (Optional) package name จะลงท้ายด้วย \_test หรือไม่ก็ได้
- เขียน test และเพื่อให้ง่ายต่อการเขียนจะใช้หลักการ AAA คือ `Arrange` `Act` และ `Assert`
  - **Arrange** สำหรับกำหนดค่าเริ่มต้นของการทดสอบ เช่น กำหนด input กำหนดผลลัพธ์ที่ต้องการ หรือ mock function ต่างๆ
  - **Act** สำหรับ execute (เรียก function) ส่วนที่ต้องการทดสอบหรือ business logic
  - **Assert** สำหรับการตรวจสอบว่า ผลการทำงานตรงตามที่คาดหวังหรือไม่ ซึ่งต้องมีทุก test

**ตัวอย่าง:**

- สมมติว่าคุณมีฟังก์ชันคำนวณผลรวมในไฟล์ `main.go`:

  ```go
  // main.go
  package main

  func Add(a, b int) int {
      return a + b
  }
  ```

- เราสามารถเขียน Unit Test สำหรับฟังก์ชันนี้ได้ดังนี้:

  ```go
  // main_test.go
  package main

  import "testing"

  func TestAdd(t *testing.T) {
      // Arrange
     input1, input2 = 4, 6
     want := 10

     // Act
     got := Add(4, 6)

     // Assert
     if got != want {
         t.Errorf("got %q, wanted %q", got, want)
     }
  }
  ```

---

### การรัน Unit Test

การรัน test ทำได้โดยใช้คำสั่ง `go test [package_name]`

```bash
go test -v gobasic
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
ok      gobasic 0.306s
```

หากการทดสอบผ่าน จะเห็นข้อความสรุปว่าการทดสอบสำเร็จ หากมีข้อผิดพลาดจะมีข้อความแจ้งเตือน

---

### การทำ Table-Driven Test

เป็นแนวทางที่นิยมในการเขียน Unit Test เพื่อทดสอบหลาย ๆ กรณีพร้อมกัน:

```go
func TestAddCases(t *testing.T) {
    cases := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"Positive numbers", 2, 3, 5},
        {"Negative numbers", -1, -1, -2},
        {"Zero", 0, 0, 0},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            result := Add(c.a, c.b)
            if result != c.expected {
                t.Errorf("Add(%d, %d) = %d; ต้องการ %d", c.a, c.b, result, c.expected)
            }
        })
    }
}
```

---

## 9. การใช้งานฟังก์ชันขั้นสูง

การใช้งานฟังก์ชันขั้นสูงใน Go มีความสำคัญเพราะช่วยให้โค้ดยืดหยุ่นและมีประสิทธิภาพมากขึ้น โดยสามารถจัดการกับพฤติกรรมต่างๆ ของฟังก์ชันได้อย่างยืดหยุ่น เช่น การใช้ค่าหลายค่าในฟังก์ชันเดียวกัน หรือการส่งฟังก์ชันเป็นอาร์กิวเมนต์ในฟังก์ชันอื่นๆ

- [Variadic Function](#variadic-function)
- [First Class Function](#first-class-function)
- [Higher Order Function](#higher-order-function)
- [Closure Function](#closure-function)

---

### Variadic Function

ฟังก์ชันแบบ Variadic ช่วยให้สามารถรับจำนวนอาร์กิวเมนต์ที่ไม่จำกัดในฟังก์ชันเดียวกันได้ โดยใช้ `...`

ซึ่งเป็นประโยชน์เมื่อเราต้องการทำงานกับชุดข้อมูลที่มีขนาดไม่แน่นอน เช่น การรวมค่าหลายค่าเข้าด้วยกัน

ตัวอย่าง:

```go
package main

import "fmt"

func sum(nums ...int) int {
 total := 0
 for _, num := range nums {
  total += num
 }
 return total
}

func main() {
 fmt.Println(sum(1, 2, 3))
 fmt.Println(sum(10, 20, 30, 40))
}

// ผลลัพธ์:
// 6
// 100
```

---

### First Class Function

ฟังก์ชันใน Go เป็น First-Class Citizen หมายความว่า ฟังก์ชันสามารถถูกส่งผ่านเป็นอาร์กิวเมนต์, คืนค่าจากฟังก์ชัน, หรือเก็บไว้ในตัวแปรได้ ซึ่งทำให้ฟังก์ชันเป็นเครื่องมือที่ทรงพลังในการออกแบบโปรแกรมที่มีความยืดหยุ่น

ตัวอย่าง:

```go
package main

import "fmt"

func add(a, b int) int {
 return a + b
}

func main() {
 fn := add             // เก็บฟังก์ชันในตัวแปร
 fmt.Println(fn(3, 4)) // เรียกใช้ฟังก์ชันผ่านตัวแปร
}

// ผลลัพธ์:
// 7
```

---

### Higher Order Function

Higher-Order Function คือฟังก์ชันที่สามารถรับฟังก์ชันอื่นเป็นอาร์กิวเมนต์ หรือส่งฟังก์ชันอื่นกลับเป็นผลลัพธ์ ฟังก์ชันประเภทนี้ช่วยให้เราสามารถสร้างลอจิกที่ยืดหยุ่นและนำกลับมาใช้ใหม่ได้ง่าย

ตัวอย่าง:

```go
package main

import "fmt"

// Higher-Order Function ที่รับฟังก์ชันเป็นพารามิเตอร์
func operate(a, b int, op func(int, int) int) int {
 return op(a, b)
}

func add(x, y int) int {
 return x + y
}

func multiply(x, y int) int {
 return x * y
}

func main() {
 fmt.Println(operate(3, 4, add))
 fmt.Println(operate(3, 4, multiply))
}

// ผลลัพธ์:
// 7
// 12
```

---

### Closure Function

Closure คือ ฟังก์ชันที่สามารถเข้าถึงตัวแปรภายนอก (จาก Scope ที่ใหญ่กว่า) และยังคงเข้าถึงตัวแปรนั้นได้แม้ฟังก์ชันนั้นจะถูกเรียกใช้นอก Scope

Closure มีประโยชน์ในการ:

- ใช้สร้างฟังก์ชันที่มี State
- ใช้ใน Higher-Order Function
- ใช้กับ Callback Function

ตัวอย่าง:

```go
package main

import "fmt"

// ฟังก์ชันที่คืนค่าฟังก์ชัน (Closure)
func counter() func() int {
    count := 0 // ตัวแปรนี้อยู่ใน Scope ของ counter()
    return func() int {
        count++ // เพิ่มค่า count ทุกครั้งที่เรียกใช้งาน
        return count
    }
}

func main() {
    c1 := counter() // สร้าง Closure ตัวแรก
    fmt.Println(c1()) // 1
    fmt.Println(c1()) // 2

    c2 := counter() // สร้าง Closure ตัวใหม่ (count เริ่มต้นใหม่)
    fmt.Println(c2()) // 1
}
```

---

## 10. การใช้ Context ใน Go

การใช้ `context` ใน Go เป็นเครื่องมือสำคัญที่ช่วยจัดการเกี่ยวกับ **Timeout**, **Deadline**, **Cancellation**, และการส่งค่าในฟังก์ชันต่างๆ โดยเฉพาะในโปรแกรมที่ต้องจัดการการทำงานแบบหลายเธรด (concurrent) หรือการทำงานที่ต้องรอผลลัพธ์ในเวลาจำกัด เช่น การทำงานกับ HTTP request หรือการประมวลผลที่ต้องทำงานร่วมกันหลายๆ ฟังก์ชัน

- [Timeout และ Deadline](#timeout-และ-deadline)
- [Cancellation](#cancellation)
- [การส่งค่า (Value Passing) ผ่าน Context](#การส่งค่า-value-passing-ผ่าน-context)

---

### Timeout และ Deadline

ในการใช้งาน `context` เพื่อจัดการ Timeout หรือ Deadline จะใช้ `context.WithTimeout` หรือ `context.WithDeadline` ซึ่งจะส่งคืน `context` ที่มีการกำหนดเวลาจำกัด (Timeout/Deadline) และสามารถใช้ในฟังก์ชันต่างๆ เช่น HTTP request หรือการทำงานที่ต้องรอผลลัพธ์จาก external service

- `context.WithTimeout` คือ ต้องการ Timeout หลังจากช่วงเวลาที่กำหนด
- `context.WithDeadline` คือ ต้องการ Timeout ที่เวลาที่กำหนด (Timestamp)

**ตัวอย่าง กำหนด Timeout ให้ฟังก์ชันหยุดทำงานหลังจากเวลาที่กำหนด:**

```go
package main

import (
 "context"
 "fmt"
 "time"
)

func longRunningTask(ctx context.Context) {
 select {
 case <-time.After(2 * time.Second): // จำลองงานที่ใช้เวลานาน
  fmt.Println("Task completed")
 case <-ctx.Done(): // เมื่อเวลา timeout หรือมีการยกเลิก
  fmt.Println("Task cancelled:", ctx.Err())
 }
}

func main() {
 // ใช้ context.WithTimeout เพื่อกำหนดเวลาสำหรับการทำงาน
 ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
 defer cancel() // ให้แน่ใจว่า cancel ถูกเรียกหลังจากการใช้งาน

 longRunningTask(ctx) // เรียกใช้งานฟังก์ชันที่ต้องรอผล
}

// ผลลัพธ์:
// Task cancelled: context deadline exceeded
```

> 💡 ฟังก์ชัน `longRunningTask` จะถูกยกเลิกทันทีที่ `context` หมดเวลา (หลังจาก 1 วินาที) ถึงแม้ว่าจะยังทำงานไม่เสร็จ

**ตัวอย่าง กำหนดให้ Context หมดเวลา ณ เวลาที่กำหนด (Timestamp):**

```go
package main

import (
 "context"
 "fmt"
 "time"
)

func longRunningTask(ctx context.Context) {
 select {
 case <-time.After(2 * time.Second): // จำลองงานที่ใช้เวลานาน
  fmt.Println("Task completed")
 case <-ctx.Done(): // เมื่อถึงเวลา deadline หรือมีการยกเลิก
  fmt.Println("Task cancelled:", ctx.Err())
 }
}

func main() {
 // กำหนด deadline เป็นเวลาปัจจุบัน + 1 วินาที
 deadline := time.Now().Add(1 * time.Second)
 ctx, cancel := context.WithDeadline(context.Background(), deadline)
 defer cancel() // ให้แน่ใจว่า cancel ถูกเรียกหลังจากการใช้งาน

 longRunningTask(ctx) // เรียกใช้งานฟังก์ชันที่ต้องรอผล
}

// ผลลัพธ์:
// Task cancelled: context deadline exceeded
```

---

### Cancellation

การยกเลิกการทำงาน (Cancellation) จะใช้เมื่อต้องการหยุดกระบวนการหรือการทำงานในกรณีที่ไม่จำเป็นต้องทำงานต่อ เช่น เมื่อผู้ใช้ปิดหน้าต่าง หรือเมื่อได้รับคำสั่งให้หยุดกระบวนการ

การยกเลิกสามารถทำได้โดยการใช้ `context.WithCancel` ซึ่งจะส่งคืน `context` ที่สามารถใช้ในการตรวจสอบการยกเลิกและฟังก์ชัน `cancel` ที่จะใช้เพื่อยกเลิกการทำงานนั้น

**ตัวอย่าง:**

```go
package main

import (
 "context"
 "fmt"
 "time"
)

func longRunningTask(ctx context.Context) {
 select {
 case <-time.After(5 * time.Second): // งานที่ใช้เวลานาน
  fmt.Println("Task completed")
 case <-ctx.Done(): // ถ้ามีการยกเลิก
  fmt.Println("Task cancelled:", ctx.Err())
 }
}

func main() {
 ctx, cancel := context.WithCancel(context.Background())
 go longRunningTask(ctx) // เรียกใช้งานงานที่ต้องใช้เวลานาน

 time.Sleep(2 * time.Second)        // จำลองการทำงานในบางช่วงเวลา
 cancel()                           // ยกเลิกการทำงานหลังจาก 2 วินาที
 time.Sleep(500 * time.Millisecond) // หยุดเพื่อให้เห็นผลลัพธ์ถูกยกเลิก
}

// ผลลัพธ์:
// Task cancelled: context canceled
```

> 💡 งานจะถูกยกเลิกเมื่อ `cancel()` ถูกเรียกใน `main` หลังจากผ่านไป 2 วินาที

---

### การส่งค่า (Value Passing) ผ่าน Context

`context` ยังสามารถใช้ในการส่งข้อมูลระหว่างฟังก์ชันต่างๆ ได้ เช่น การส่งค่าที่ต้องการใช้ในหลายๆ ฟังก์ชันหรือหลายๆ งาน โดยใช้ `context.WithValue` ซึ่งจะส่งคืน `context` ที่มีข้อมูลบางอย่างที่เกี่ยวข้องกับกระบวนการนั้น

**ตัวอย่าง:**

```go
package main

import (
 "context"
 "fmt"
)

func process(ctx context.Context) {
 // ตรวหาค่าใน context ด้วย Key "key"
 if value := ctx.Value("key"); value != nil {
  fmt.Println("Received value:", value) // ถ้ามีค่า, แสดงผลลัพธ์
 } else {
  fmt.Println("No value found")
 }
}

func main() {
 // สร้าง context ที่มีค่า "key" คือ "myValue"
 ctx := context.WithValue(context.Background(), "key", "myValue")
 process(ctx) // ส่งค่า "myValue" ผ่าน context
}

// ผลลัพธ์:
// Received value: myValue
```

> 💡 ข้อมูล (`"myValue"`) จะถูกส่งผ่าน `context` และสามารถนำไปใช้ในฟังก์ชันที่ต้องการได้

---

## 11. การทำงานกับ JSON

การทำงานกับ **JSON** ใน Go ใช้แพ็กเกจ `encoding/json` ที่มาพร้อมกับ Go ซึ่งมีฟังก์ชันสำหรับการแปลงข้อมูลระหว่าง **Go struct** และ **JSON** format (ทั้งการแปลงจาก struct ไปเป็น JSON และจาก JSON ไปเป็น struct)

- [การแปลงจาก struct เป็น JSON (Marshal)](#การแปลงจาก-struct-เป็น-json-marshal)
- [การแปลงจาก JSON เป็น struct (Unmarshal)](#การแปลงจาก-json-เป็น-struct-unmarshal)
- [การจัดการกับ JSON ที่มี Key ไม่ตรงกับฟิลด์ใน struct](#การจัดการกับ-json-ที่มี-key-ไม่ตรงกับฟิลด์ใน-struct)

---

### การแปลงจาก struct เป็น JSON (Marshal)

การแปลงจาก `struct` ไปเป็น JSON สามารถทำได้โดยใช้ฟังก์ชัน `json.Marshal` ซึ่งจะรับ `struct` เป็นอาร์กิวเมนต์และคืนค่าเป็น JSON ในรูปแบบของ **byte slice** (เช่น `[]byte`)

**ตัวอย่าง:**

```go
package main

import (
 "encoding/json"
 "fmt"
)

// สร้าง struct ที่จะใช้ในการแปลง
type Person struct {
 Name    string
 Age     int
 Address string
}

func main() {
 person := Person{
  Name:    "John Doe",
  Age:     30,
  Address: "123 Main St",
 }

 // แปลง struct ไปเป็น JSON
 jsonData, err := json.Marshal(person)
 if err != nil {
  fmt.Println("Error marshaling:", err)
  return
 }

 // แสดงผล JSON ในรูปแบบ string
 fmt.Println(string(jsonData))
}

// ผลลัพธ์:
// {"Name":"John Doe","Age":30,"Address":"123 Main St"}
```

---

### การแปลงจาก JSON เป็น struct (Unmarshal)

การแปลงจาก JSON ไปเป็น `struct` ใช้ฟังก์ชัน `json.Unmarshal` ซึ่งรับ JSON ในรูปแบบของ **byte slice** (เช่น `[]byte`) และแปลงมันเป็น `struct` ที่กำหนด

**ตัวอย่าง:**

```go
package main

import (
 "encoding/json"
 "fmt"
)

// สร้าง struct ที่ใช้ในการแปลง JSON
type Person struct {
 Name    string
 Age     int
 Address string
}

func main() {
 jsonData := `{"Name":"John Doe","Age":30,"Address":"123 Main St"}`

 var person Person

 // แปลง JSON ไปเป็น struct
 err := json.Unmarshal([]byte(jsonData), &person)
 if err != nil {
  fmt.Println("Error unmarshaling:", err)
  return
 }

 // แสดงผล struct
 fmt.Println(person)
}

// ผลลัพธ์:
// {John Doe 30 123 Main St}
```

---

### การจัดการกับ JSON ที่มี Key ไม่ตรงกับฟิลด์ใน struct

ในบางกรณี JSON ที่ได้รับอาจมีคีย์ที่ไม่ตรงกับฟิลด์ใน `struct` หรือไม่ต้องการให้แปลงบางฟิลด์ใน `struct` ไปเป็น JSON เราสามารถใช้ JSON Tag ในการควบคุมการแปลงได้

- ใช้ `json:"fieldname"` เพื่อกำหนดชื่อฟิลด์ใน JSON
- ใช้ `json: "omitempty"` เพื่อไม่ให้ฟิลด์ที่มีค่าเป็น zero-value (เช่น `""` หรือ `0`) ปรากฏใน JSON
- ใช้ `json:"-"` เพื่อไม่ให้ฟิลด์นั้นถูกแปลงเป็น JSON หรือไม่ให้รับค่าเมื่อแปลงจาก JSON

**ตัวอย่าง:**

```go
package main

import (
 "encoding/json"
 "fmt"
)

type Person struct {
 Name    string `json:"name"`
 Age     int    `json:"age,omitempty"` // กรณีที่เป็น 0 จะไม่แสดงใน JSON
 Address string `json:"-"`             // ไม่ให้แปลง Address ไปเป็น JSON
}

func main() {
 person := Person{
  Name:    "John Doe",
  Address: "123 Main St",
 }

 jsonData, _ := json.Marshal(person)
 fmt.Println(string(jsonData))
}

// ผลลัพธ์:
// {"name":"John Doe"}
```

---

## สรุป

ภาษา Go เป็นภาษาที่มีประสิทธิภาพสูง ใช้งานง่าย และเหมาะสำหรับการพัฒนาแอปพลิเคชันที่ต้องการความเร็วและรองรับการประมวลผลพร้อมกัน บทความนี้ได้ครอบคลุมเนื้อหาตั้งแต่การเริ่มต้นใช้งาน พื้นฐานของภาษา ไปจนถึงแนวคิดขั้นสูง เช่น Interface, Generics, Concurrency และการจัดการข้อผิดพลาด

เมื่อคุณเข้าใจหลักการพื้นฐานและแนวคิดสำคัญเหล่านี้แล้ว คุณจะสามารถนำความรู้ไปประยุกต์ใช้ในการพัฒนาแอปพลิเคชันจริงได้อย่างมั่นใจ และสามารถขยายขอบเขตความสามารถของคุณไปสู่การสร้างระบบที่มีความซับซ้อนมากยิ่งขึ้น

การเรียนรู้ภาษา Go เป็นการลงทุนที่คุ้มค่า ไม่ว่าคุณจะเป็นนักพัฒนามือใหม่หรือมีประสบการณ์แล้ว การมีความเชี่ยวชาญในภาษา Go จะช่วยให้คุณก้าวทันเทคโนโลยีและสามารถพัฒนาโซลูชันที่มีประสิทธิภาพในโลกของซอฟต์แวร์ยุคปัจจุบัน
