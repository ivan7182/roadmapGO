package main

import (
	"fmt"
)

/*
Soal 1: Membuat Tipe Data Baru
Buatlah program Golang yang mendeklarasikan tipe data baru bernama Centimeter
yang merupakan tipe alias dari tipe int. Gunakan tipe ini untuk menghitung
jumlah total panjang dari tiga nilai berikut:

Nilai pertama: 150 cm
Nilai kedua: 200 cm
Nilai ketiga: 175 cm

no 2
Buat tipe data baru bernama Kilogram yang merupakan
tipe alias dari float64. Buatlah program untuk mengonversi berat
badan dalam kilogram ke gram menggunakan rumus:

1 kilogram = 1000 gram
Gunakan nilai berat berikut:

Berat pertama: 62.5 kg
Berat kedua: 70.3 kg
*/
func main() {
	// tipedatabaru()
	conversi()

}

func tipedatabaru() {
	type centimeter int

	var nilaiPertama centimeter = 150
	var nilaiKedua centimeter = 200
	var nilaiKetiga centimeter = 175

	totalPanjang := nilaiPertama + nilaiKedua + nilaiKetiga

	fmt.Printf("total panjang adalah %d cm.\n", totalPanjang)

}

func conversi() {
	type kilogram float64

	var beratPertama kilogram = 62.5
	var beratKedua kilogram = 70.3

	totalBerat := (float64(beratPertama) + float64(beratKedua)) * 100
	fmt.Printf("total berat dalam gram adalah %.0f gram.\n", totalBerat)

}
