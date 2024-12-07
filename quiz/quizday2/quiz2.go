package main

import "fmt"

func main() {
	/*
		1. (fizz buzz)buatkan angka 1 - 100
		namun setiap angka yang habis dibagi 3 cetak feez
		dan  dibagi 5 buzz, jika habis dibagi 3 dan 5 fezz buzz
		2. buatkan operator matematika sederhana
	*/

	// fizzbuzz()
	calculate()

}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func calculate() {
	var operator string
	fmt.Println("masukan operator : +, _, *, /")
	fmt.Scan(&operator)

	var angka1, angka2 float64
	fmt.Println("masukan angka :")
	fmt.Scan(&angka1, &angka2)

	switch operator {
	case "+":
		fmt.Println(angka1 + angka2)
	case "-":
		fmt.Println(angka1 - angka2)
	case "/":
		fmt.Println(angka1 / angka2)
	case "*":
		fmt.Println(angka1 * angka2)
	default:
		{
			fmt.Println("eror")
		}
	}
}
