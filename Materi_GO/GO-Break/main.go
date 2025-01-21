package main

import "fmt"

func main() {
	// breakk()
	continuee()
}

func breakk() {
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("perulangan ke", i)
	}
}

func continuee() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan ke", i)
	}
}
