package main

import "fmt"

func main() {
	var a uint8 = 9
	var w, x, y, z uint8 = 5, 10, 15, 20

	w %= 5 // w = w + 5
	x += 9 // x = 10 + 9
	y *= 10
	z /= 5

	fmt.Println("niai a :", a)
	fmt.Println("niai w :", w)
	fmt.Println("niai x :", x)
	fmt.Println("niai y :", y)
	fmt.Println("niai z :", z)
}
