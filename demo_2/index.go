package main

import (
	"fmt"
	"time"
)

func main() {
	var msg string = "Hello GO from msg"
	fmt.Println("Hello GO!" + msg)
	time.Sleep(6 * time.Second)
	fmt.Println("Hello GO!2" + msg)
}
