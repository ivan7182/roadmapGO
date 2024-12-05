package main

import "fmt"

func main() {
	// a := 19

	// fmt.Println(a > 10 && a < 8)   // false
	// fmt.Println(a < 7 || a > 9)    // true
	// fmt.Println(!(a < 7 || a > 9)) // false

	nilaiUjian := 80
	nilaiAbsen := 85

	lulusUjian := nilaiUjian > 75 //true
	lulusAbsen := nilaiAbsen < 90 // true

	lulus := lulusAbsen && lulusUjian //true
	fmt.Println(lulus)
}
