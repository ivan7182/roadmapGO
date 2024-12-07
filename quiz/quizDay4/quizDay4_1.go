package main

import "fmt"

/*
1. Membuat Custom Type
Buat sebuah custom type bernama Age yang merupakan representasi
dari tipe int. Lalu, buat variabel
myAge bertipe Age dengan nilai 25. Cetak nilai dari variabel tersebut.
2. Custom Type untuk Boolean
Buat sebuah custom type bernama IsActive yang
merupakan representasi dari tipe bool. Lalu, buat variabel status bertipe IsActive
dengan nilai true. Cetak nilai dari variabel tersebut.
3. Menggunakan Type Alias
Gunakan deklarasi type alias untuk mendefinisikan tipe ID sebagai
alias dari string. Lalu buat variabel userID bertipe ID dengan nilai "U12345".
Cetak nilai dari variabel tersebut.
4.  Array dengan Custom Type
Buat sebuah custom type bernama Scores yang merupakan representasi dari
array int dengan panjang 3. Buat variabel studentScores bertipe Scores dengan nilai [85, 90, 78].
Cetak nilai dari variabel tersebut.
*/

func main() {
	// soal1()
	// soal2()
	// soal3()
	soal4()

}

func soal1() {
	type age int

	var myAge age = 25
	fmt.Println(myAge)
}

func soal2() {
	type IsActive bool

	var status IsActive = true
	fmt.Println(status)
}

func soal3() {
	type ID string

	var userID ID = "U12345"

	fmt.Println(userID)
}

func soal4() {
	type scores []int

	var studentScores scores = []int{85, 90, 78}
	fmt.Println(studentScores)

}
