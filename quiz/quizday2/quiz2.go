package main

import "fmt"

func main() {
	/*
		1. (feez buzz)buatkan angka 1 - 100
		namun setiap angka yang habis dibagi 3 cetak feez
		dan  dibagi 5 buzz, jika habis dibagi 3 dan 5 fezz buzz
		2. buatkan operator matematika
	*/
	// FizzBuzz()
	calculateInput()
}

func FizzBuzz() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func calculate(a int, b int, operator string) int {
	var hasil int
	switch operator {
	case "+":
		hasil = a + b
	case "-":
		hasil = a - b
	case "*":
		hasil = a * b
	case "/":
		hasil = a / b

	}
	return hasil

}

func calculateInput() {
	var a, b int
	var operator string
	fmt.Println("input angka 1:")
	fmt.Scan(&a)
	fmt.Println("input angka 1:")
	fmt.Scan(&b)
	fmt.Println("InputOperator :")
	fmt.Scan(&operator)
	fmt.Println("hasil : ", calculate(a, b, operator))
}
