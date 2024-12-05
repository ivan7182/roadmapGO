package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("masukan nilai : ")
	fmt.Scan(&nilai)

	//if else
	// if nilai == 100 {
	// 	fmt.Println("A")
	// } else if nilai < 100 && nilai > 75 {
	// 	fmt.Println("B")
	// } else if nilai <= 75 && nilai > 40 {
	// 	fmt.Println("C")
	// } else {
	// 	fmt.Println("remidi")
	// }

	switch {
	case nilai == 100:
		fmt.Println("A")
	case (nilai < 75) && (nilai > 40):
		fmt.Println("B")
	case (nilai <= 40) && (nilai > 10):
		fmt.Println("C")
	default:
		fmt.Println("remidi")
	}

}
