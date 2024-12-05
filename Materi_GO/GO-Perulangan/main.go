package main

import "fmt"

func main() {
	// Keyword for 1
	// for i := 5; i < 9; i++ {
	// 	fmt.Println("angka", i)
	// }

	// Keyword for 2
	// var i = 3
	// for i <= 9 {
	// 	fmt.Println("angka", i)
	// 	i++
	// }

	// for menggunakan break
	var i = 4
	for {
		fmt.Println("angka", i)

		i++

		if i >= 9 {
			break
		}
	}
}
