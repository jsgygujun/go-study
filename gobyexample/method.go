package main

import "fmt"

type Rect struct {
	width, height int
}

// 接收 Rect 指针
func (r *Rect) area() int { // 这里的 area 方法有一个接收器类型 rect。
	return r.width * r.height
}

// 接收 Rect 值
func (r Rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := Rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area()) // Go 自动处理方法调用时的值和指针之间的转化
	fmt.Println("perim", rp.perim())
}
