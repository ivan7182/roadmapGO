package main

import "fmt"

func greet(name string) (text string) {
	text = "hallo " + name
	return
}

func add(x, y int) int {
	return x + y
}

func subAdd(x, y int) (addRest int, subRest int) {
	return x + y, x - y
}

func main() {
	// fmt.Println(greet("ivan"))
	fmt.Println(add(2, 7))
	fmt.Println(subAdd(2, 7))
}
