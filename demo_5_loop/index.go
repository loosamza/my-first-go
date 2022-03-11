package main

import "fmt"

func main() {
	fnFor()
	fnWhile()
	fnForEach()
	fnWHileUsingBreak()
}

func fnFor() {
	for index := 0; index <= 10; index++ {
		fmt.Printf("Index %d\n", index)
	}
}

func fnWhile() {
	index := 0
	for index < 5 {
		fmt.Printf("Index %d\n", index)
		index++
	}
}

func fnForEach() {
	todoList := []string{"Do it", "Do it more", "Do it more more"}
	for index, item := range todoList {
		fmt.Printf("%d. %s\n", index+1, item)
	}
	// ignore index
	for _, item := range todoList {
		fmt.Printf("%s\n", item)
	}
}

func fnWHileUsingBreak() {
	index := 0
	for true {
		if index > 5 {
			break
		}
		fmt.Printf("Before break index %d\n", index)
		index++
	}
}
