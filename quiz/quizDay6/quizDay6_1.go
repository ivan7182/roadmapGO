package main

import (
	"fmt"
)

func main() {
	//soal1()
	//soal3()
	//soal4()
	soal5()
}

func soal1() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}

	arr3 := append(arr1, arr2...)
	fmt.Println(arr3)

}

func soal3() {
	arr := []int{10, 20, 30, 40, 50}

	fmt.Println(len(arr))
}

func soal4() {
	arr := []string{"apple", "banana", "cherry"}
	fmt.Println(arr)
}

func soal5() {
	arr := []int{1, 2, 3, 4, 5}
	var element int
	fmt.Println("masukan element yang dicar :")
	fmt.Scan(&element)

	found := false
	for _, num := range arr {
		if num == element {
			found = true
			break
		}
	}
	if found {
		fmt.Println("ada")
	} else {
		fmt.Println("tidak ada")
	}
}
