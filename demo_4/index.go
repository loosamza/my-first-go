package main

import (
	"fmt"
)

func main() {
	fn1_no_arg()
	fn2_arg("show arg")
	fn3("Print count ", 1)

	const a, b = 2, 3
	fmt.Printf("%d+%d=%d\n", a, b, sum(a, b))

	x, y := swap(a, b)
	fmt.Printf("%d,%d => %d,%d\n", a, b, x, y)

	_x, _y, title := swap2(10, 20, "NAME")
	fmt.Printf("%d,%d,%s => %d,%d,%s\n", 10, 20, "NAME", _x, _y, title)
}

func fn1_no_arg() {
	fmt.Println("SSS")
}

func fn2_arg(msg string) {
	fmt.Println(msg)
}

func fn3(title string, version int) {
	fmt.Print(title)
	fmt.Println(version)
}

// single return
func sum(a int, b int) int {
	return a + b
}

//mutiple return
func swap(a int, b int) (int, int) {
	return b, a
}

//mutiple return with name
func swap2(a int, b int, c string) (x, y int, title string) {
	y = a
	x = b
	title = c
	return
}
