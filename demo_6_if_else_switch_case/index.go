package main

import "fmt"

func main() {
	someValue := 11
	if someValue == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!= 10")
	}

	if someValue > 10 && someValue < 20 {
		fmt.Println("someValue > 10 && someValue < 20")
	} else {
		fmt.Println("not someValue > 10 && someValue < 20")
	}
	if result := doSomething(); result == "ok" {
		fmt.Println("Ok")
	} else {
		fmt.Println("Not Ok")
	}

	fnSwitchCase()
}

func doSomething() string {
	//dosomething
	return "ok"
}

func fnSwitchCase() {
	index := 2
	switch index {
	case 0:
		fmt.Println(0)
		break
	case 1:
		fmt.Println(2)
		break
	case 2:
		fmt.Println(2)
		break
	default:
		fmt.Println("Something else")
		break
	}
}
