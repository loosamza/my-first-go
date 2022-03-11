package main

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

}
