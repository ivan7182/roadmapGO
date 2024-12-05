package main

import "fmt"

func main() {
	a := [3]int{100, 200, 3000}

	var b = [...]int{1, 2, 3, 4, 5}

	var pemrograman = [2]string{"Go", "Java"}

	harga := [3]int{100, 300, 400}
	harga[0] = 900

	// //iinisialisai
	number := [2]int{}

	//perulangan elemen array
	var cars = [3]string{"bmw", "ferarri", "toyota"}
	for i := 0; i < len(cars); i++ {
		fmt.Printf("elemen %d : %s\n", i, cars[i])
	}

	// //for - range
	// var sembako = [3]string{"gula", "terigu", "beras"}
	// for _, sembak := range sembako {
	// 	fmt.Printf("sembako : %s\n", sembak)
	// }

	//alokasikan keyword make
	var sembako = make([]string, 4)
	sembako[0] = "gula"
	sembako[1] = "terigu"
	sembako[2] = "beras"
	sembako[3] = "kecap"

	fmt.Println(sembako)

	fmt.Println(a)
	fmt.Println(len(b))
	fmt.Println(len(pemrograman))
	fmt.Println(harga)
	fmt.Println(number)
}
