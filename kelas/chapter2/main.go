package main

import "fmt"

func greet(name string) (text string) {
	text = "hallo" + name
	return
}

func add(x, y int) int {
	return x + y
}

func addSub(x, y int) (addRes int, subRes int) {
	return x + y, x - y
}

func main() {

	fmt.Println(add(2, 5))
	fmt.Println(greet("ivan"))
	fmt.Println(addSub(2, 5))
}
