package main

import "fmt"

/*
soal 1:
Buatlah program yang mencetak angka dari 1 hingga 10, tetapi berhenti
mencetak jika menemukan angka yang habis dibagi 4.
Berapa angka terakhir yang dicetak?

soal 2:
Buatlah program yang mencetak angka dari 1 hingga 15,
tetapi melewatkan angka yang merupakan kelipatan 5.

Berapa angka yang tidak dicetak?
*/
func main() {
	// soal1()
	soal2()
}

func soal1() {
	for i := 1; i < 10; i++ {
		if i%4 == 0 {
			break
		}

		fmt.Println("perulangan ke", i)
	}
}

func soal2() {
	for i := 1; i < 15; i++ {
		if i%5 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
