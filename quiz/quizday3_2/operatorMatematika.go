package main

import "fmt"

/*
Soal 1: Menghitung Luas dan Keliling Persegi Panjang
Buatlah program yang meminta pengguna memasukkan panjang dan lebar sebuah
persegi panjang, lalu menghitung dan menampilkan:
Luas persegi panjang.
Keliling persegi panjang.
Rumus:
Luas = panjang × lebar
Keliling = 2 × (panjang + lebar)

soal 2: Menentukan Ganjil atau Genap
Buatlah program yang meminta pengguna memasukkan
sebuah angka, kemudian menentukan apakah angka tersebut adalah bilangan ganjil atau genap.

soal 3: Menampilkan Angka Kelipatan Tertentu
Buatlah program yang meminta pengguna memasukkan dua angka:
Batas atas (misalnya 20)
Bilangan kelipatan (misalnya 5)
Program akan menampilkan semua angka kelipatan bilangan tersebut hingga batas atas.
Contoh:
Batas atas: 20
Bilangan kelipatan: 5
Output: 5, 10, 15, 20

soal 4 : Menghitung Nilai Akhir Siswa
Buatlah program yang meminta pengguna untuk memasukkan nilai ujian
dan nilai tugas siswa. Program akan menghitung nilai akhir dengan rumus:
Nilai Akhir = (Nilai Ujian * 70%) + (Nilai Tugas * 30%).
Kemudian, program akan menampilkan nilai akhir dan memberi pesan lulus jika
nilai akhir lebih dari atau sama dengan 70, atau tidak lulus jika kurang dari 70.


*/
func main() {
	// soal1()
	// soal2()
	// soal3()
	soal4()
}

func soal1() {
	var panjang int
	fmt.Println("masukan panjang Persegi panjang :")
	fmt.Scan(&panjang)

	var lebar int
	fmt.Println("masukan lebar persegi panjang :")
	fmt.Scan(&lebar)

	luas := panjang * lebar
	keliling := (panjang + lebar) * 2

	fmt.Printf("luasnya adalah : %d\n", luas)
	fmt.Printf("kelilingnya adalah :  %d\n", keliling)
}

func soal2() {
	var angka int
	fmt.Println("masukan angka :")
	fmt.Scan(&angka)

	if angka%2 == 0 {
		fmt.Println("bilangan genap")
	} else {
		fmt.Println("Bilangan ganjil")
	}
}

func soal3() {
	var angka int
	fmt.Println("masukan angka :")
	fmt.Scan(&angka)

	var batas int
	fmt.Println("masukan batas angka :")
	fmt.Scan(&batas)

	fmt.Printf("angka kelipatan %d hingga %d adalah : \n", angka, batas)

	for i := angka; i <= batas; i += angka {
		fmt.Println(i)
	}
}

func soal4() {
	var nilaiUjian float64
	fmt.Println("masukan nilai Ujian :")
	fmt.Scan(&nilaiUjian)

	var nilaiTugas float64
	fmt.Println("masukan nilai tugas :")
	fmt.Scan(&nilaiTugas)

	var nilaiAkhir = (nilaiTugas * 0.7) + (nilaiUjian * 0.3)

	fmt.Printf("nilai akhir siswa %.2f\n", nilaiAkhir)
	if nilaiAkhir >= 70 {
		fmt.Println("lulus")
	} else {
		fmt.Println("tidak lulus")
	}
}
