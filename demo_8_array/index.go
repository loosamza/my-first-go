package main

import "fmt"

func main() {
	var array1 []int = []int{1, 2, 3, 4}
	var array2 = []int{1, 2, 3, 4}
	var array3 [3]string
	for _, item := range array1 {
		print(item, ",")
	}
	for _, item := range array2 {
		print(item, ",")
	}
	array3[0], array3[1], array3[2] = "android", "ios", "nokia"
	for _, item := range array3 {
		print(item, ",")
	}
	println()
	var numbers1 = make([]int, 3, 5)
	var numbers2 []int
	numbers1 = append(numbers1, 1)
	numbers1 = append(numbers1, 2)
	numbers1 = append(numbers1, 3)
	showSlice(numbers1, "numbers1")
	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	showSlice(numbers2, "numbers2")

	var removeArray1 = []int{1, 2, 3, 4, 5, 6}

	//leading remove
	showSlice(removeArray1, "Full array ")
	removeArray1 = removeArray1[1:len(removeArray1)]
	showSlice(removeArray1, "leading remove ")
	removeArray1 = removeArray1[1:len(removeArray1)]
	showSlice(removeArray1, "leading remove ")

	// trailing remove
	removeArray1 = removeArray1[0 : len(removeArray1)-1]
	showSlice(removeArray1, "trailing remove")

	// remove specific index
	removeArray1 = removeIndex(removeArray1, 2)
	showSlice(removeArray1, "remove specific index")
}

func showSlice(x []int, s string) {
	fmt.Printf("%s len=%d cap=%d slice=%v\n", s, len(x), cap(x), x)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
