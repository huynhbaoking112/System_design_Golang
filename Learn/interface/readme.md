# Interface trong Go

## Giới thiệu
Trong Go, **interface** là một tập hợp các phương thức mà một kiểu dữ liệu (struct) có thể triển khai. Interface giúp đạt được tính **đa hình** (polymorphism) bằng cách cho phép các kiểu dữ liệu khác nhau thực thi cùng một tập hợp phương thức.

---

## 1. Định nghĩa Interface
### Ví dụ: Interface `Shape`
```go
package main

import (
	"fmt"
)

// Định nghĩa interface Shape
type Shape interface {
	Area() float64
}

// Struct Circle
type Circle struct {
	Radius float64
}

// Struct Rectangle
type Rectangle struct {
	Width, Height float64
}

// Implement phương thức Area() cho Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Implement phương thức Area() cho Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Khai báo một biến kiểu Shape
	var s Shape

	// Gán một Circle vào s
	s = Circle{Radius: 5}
	fmt.Println("Circle Area:", s.Area())

	// Gán một Rectangle vào s
	s = Rectangle{Width: 4, Height: 6}
	fmt.Println("Rectangle Area:", s.Area())
}
```

### Giải thích:
- Interface `Shape` có một phương thức `Area() float64`.
- `Circle` và `Rectangle` đều có phương thức `Area()`, vì vậy chúng **triển khai** interface `Shape`.
- Biến `s` có kiểu `Shape`, nhưng có thể chứa **bất kỳ struct nào triển khai `Shape`**.
- Khi gọi `s.Area()`, Go sẽ gọi đúng phương thức dựa trên kiểu thực tế (`Circle` hoặc `Rectangle`).

---

## 2. Interface rỗng (`interface{}`)
Trong Go, `interface{}` là một **interface rỗng**, có thể chứa bất kỳ kiểu dữ liệu nào.

### Ví dụ:
```go
func PrintValue(value interface{}) {
	fmt.Println("Value:", value)
}

func main() {
	PrintValue(42)       // int
	PrintValue("Hello")  // string
	PrintValue(3.14)     // float64
}
```
### Công dụng:
- **`interface{}`** hữu ích khi bạn muốn viết hàm có thể nhận **bất kỳ kiểu dữ liệu nào**.

---

## 3. Kiểm tra kiểu dữ liệu với `type assertion`
Khi sử dụng interface rỗng (`interface{}`), bạn có thể kiểm tra kiểu dữ liệu thực tế bằng `type assertion`.

### Ví dụ:
```go
func CheckType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	CheckType(100)
	CheckType("GoLang")
	CheckType(3.14)
}
```

---

## Kết luận
✅ Interface trong Go giúp định nghĩa **tập hợp phương thức** mà một kiểu dữ liệu có thể triển khai.

✅ Interface hỗ trợ **đa hình** bằng cách cho phép một biến có thể chứa nhiều kiểu khác nhau.

✅ `interface{}` là interface đặc biệt có thể chứa mọi kiểu dữ liệu.

✅ Có thể dùng **type assertion** để kiểm tra kiểu thực tế khi dùng interface.

⏩ **Go sử dụng interface để viết code linh hoạt và dễ mở rộng hơn!** 🚀

