package main

import "fmt"

func main() {
	msg := "Some message"
	var msgPointer *string = &msg

	fmt.Println(msg)
	fmt.Println(msgPointer)
	fmt.Println(*msgPointer)
	changeMsg(&msg, "New message 1")
	fmt.Println(msg)
	fmt.Println(msgPointer)
	fmt.Println(*msgPointer)
	changeMsg(msgPointer, "New message 2")
	fmt.Println(msg)
	fmt.Println(msgPointer)
	fmt.Println(*msgPointer)
}

func changeMsg(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
