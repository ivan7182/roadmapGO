package main

import "fmt"

/*
Soal 1: Fungsi Sederhana Buatlah sebuah fungsi dalam
Golang yang menerima dua parameter bilangan bulat dan mengembalikan
hasil penjumlahan dari kedua bilangan tersebut. Berikan contoh penggunaannya.
*/

func main() {
	angka1 := 10
	angka2 := 40

	hasil := tambahh(angka1, angka2)
	fmt.Printf("hasil penjumlahan %d + %d = %d\n", angka1, angka2, hasil)

}

func tambahh(a int, b int) int {
	return a + b
}
