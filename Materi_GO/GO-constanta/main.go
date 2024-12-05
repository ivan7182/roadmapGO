package main

import "fmt"

const LENGHT int = 9 // bisa di deklarasikan diluar maupun didalam function
const WIDTH = 2

// Multiple Const 1
const (
	X              = 10
	ISTODAY  bool  = true
	NUMERIC  uint8 = 1
	FLOATNUM       = 2.2
)

// Deklarasi Multipe Const 2
const (
	TODAY = "senin"
	SEKARANG
	/* nilai tidak dituliskan dalam deklarasi
	konstanta, maka tipe data dan nilai yang dipergunakan adalah
	sam a seperti konstanta yang dideklarasikan diatasnya. */
)

func main() {

	// const PI = 10
	// // PI = 12 // tidak dapat diubah valuenya

	//deklrasi multiple konstanta dalam satu baris
	const SENIN, SELASA = 1, 2
	const TIGA, EMPAT = "tiga", "empat"

	fmt.Println(LENGHT)
	fmt.Println(WIDTH)
	// fmt.Println(PI)

	fmt.Println(X)
	fmt.Println(ISTODAY)
	fmt.Println(NUMERIC)
	fmt.Println(FLOATNUM)

	fmt.Println(TODAY)
	fmt.Println(SEKARANG)

	fmt.Println(SENIN)
	fmt.Println(SELASA)
	fmt.Println(TIGA)
	fmt.Println(EMPAT)

}
