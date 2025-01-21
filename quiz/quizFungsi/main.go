package main

import "fmt"

/*
Soal 1: Fungsi Sederhana Buatlah sebuah fungsi dalam
Golang yang menerima dua parameter bilangan bulat dan mengembalikan
hasil penjumlahan dari kedua bilangan tersebut. Berikan contoh penggunaannya.
Soal 2: Fungsi Genap Ganjil Tuliskan sebuah fungsi dalam Golang untuk mengecek apakah sebuah bilangan bulat
positif yang diberikan sebagai input adalah genap atau ganjil. Fungsi ini harus mengembalikan string "Genap"
jika bilangan tersebut genap, dan "Ganjil" jika bilangan tersebut ganjil.
*/

func main() {
	// angka1 := 10
	// angka2 := 40

	// hasil := tambahh(angka1, angka2)
	// fmt.Printf("hasil penjumlahan %d + %d = %d\n", angka1, angka2, hasil)

	angka := 19
	hasil := cekGenapGanjil(angka)
	fmt.Printf("angka %d adalah %s\n", angka, hasil)
}

// func tambahh(a int, b int) int {
// 	return a + b
// }

func cekGenapGanjil(angka int) string {
	if angka%2 == 0 {
		return "Genap"
	}
	return "Ganjil"
}
