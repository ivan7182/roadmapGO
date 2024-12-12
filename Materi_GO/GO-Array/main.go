package main

import "fmt"

func main() {
	//arrayMultidimensi()
	// arrayFor()
	// arrayRange()
	//_for()
	makeArray()

}

func arrayMultidimensi() {
	var sembako = [2][3]string{{"beras", "gula", "terigu"}, {"telor", "minyak", "garam"}}
	fmt.Println(sembako)
	// fmt.Println(number2)
}

func arrayFor() {
	var number = [2]int{1, 2}

	for i := 0; i < len(number); i++ {
		fmt.Printf("elemen %d : %d\n", i, number[i])
	}
}

func arrayRange() {
	var sembako = [2]string{"beras", "gula"}

	for i, sembakoSaya := range sembako {
		fmt.Println(i, sembakoSaya)
	}
}

func _for() {
	var sembako = [2]string{"beras", "gula"}

	for _, sembakoSaya := range sembako {
		fmt.Println(sembakoSaya)
	}
}

func makeArray() {

	var sembako = make([]string, 3)

	sembako[0] = "gula"
	sembako[1] = "beras"
	sembako[2] = "terigu"

	fmt.Println(sembako)
}
