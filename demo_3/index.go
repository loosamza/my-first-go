package main

import "fmt"

var count = 0

func main() {
	fmt.Println("Begin")

	// Explicit Declaration
	var tmp1 int = 0
	tmp1 = 10
	var tmp2 string = "String"
	var tmp3 bool = true
	const tmp4 = 5
	// cannot assign to tmp4 tmp4 = 6

	// Implicait Declaratin
	tmp5 := 10

	fmt.Println(tmp1)
	count++
	fmt.Println(tmp2)
	count++
	fmt.Println(tmp3)
	count++
	fmt.Println(tmp4)
	count++
	fmt.Println(tmp5)
	count++
	run()
}

func run() {
	fmt.Println(count)
}
