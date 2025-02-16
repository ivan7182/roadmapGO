package main

import "fmt"

func print_hello() {
	fmt.Printf("hello\n")
}
func main() {

	//defer
	defer print_hello()
	// if, else if, else
	// age := 29

	// if age < 18 {
	// 	fmt.Println("Belum cukup umur")
	// } else if age > 60 {
	// 	fmt.Println("Terlalu Tua")
	// } else {
	// 	fmt.Println("selamat datang")
	// }

	//switch

	// nilai := "G"

	// switch nilai {
	// case "A":
	// 	fmt.Println("sangatt baik")
	// case "C", "B":
	// 	fmt.Println("baik")
	// 	fmt.Println("tingkatkan lagi")
	// case "E":
	// 	fmt.Println("temui guru")
	// 	fallthrough
	// default:
	// 	fmt.Println("remidi")
	// }

	//Looping
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("looping %d\n", i)
	// }

	// i := 0
	// for {
	// 	i += 1
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Printf("looping %d\n", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	numbers := []int64{1, 2, 3, 4, 5}
	for i := range numbers {
		fmt.Printf("saya mendapatkan %d di indeks %d\n", numbers[i], i)
	}
}
