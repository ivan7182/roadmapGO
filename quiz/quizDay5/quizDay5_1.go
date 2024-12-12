package main

import "fmt"

func main() {
	soal1()
}

func soal1() {
	var number = [5]int{1, 2, 3, 4, 20}
	var sum int
	fmt.Println(len(number))

	for _, value := range number {
		sum += value
	}

	fmt.Println("jumlah element array", sum)

}
