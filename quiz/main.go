package main

import (
	"fmt"
	"math"
)

func main() {
	hitungLuasLingkaran()
	cetakBilanganGenap()
	menukarValue()
}

// 1. hitung luas lingkaran
func hitungLuasLingkaran() {
	var radius float64

	fmt.Print("masukan jari - jari lingkaran :")
	fmt.Scan(&radius)

	area := math.Pi * math.Pow(radius, 2)
	lingkaran := 2 * math.Pi * radius

	fmt.Printf("area lingkaran : %.2f\n", area)
	fmt.Printf("luas lingkaran : %.2f\n", lingkaran)
}

// 2. cetak bilangan genap
func cetakBilanganGenap() {
	fmt.Printf("Deret Bilangan Genap 1 - 20")
	for i := 1; i <= 20; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

// 3. menukar pointer

func swap(a, b *int) {
	*a, *b = *b, *a
}

func menukarValue() {
	var x, y int

	fmt.Print("masukan nilai x :")
	fmt.Scan(&x)

	fmt.Print("masukan nilai y :")
	fmt.Scan(&y)

	fmt.Printf("sebelum pertukaran: x = %d, y = %d\n", x, y)

	swap(&x, &y)

	fmt.Printf("sesudah pertukaran : x = %d, y = %d\n", x, y)

}
