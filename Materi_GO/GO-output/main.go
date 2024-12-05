package main

import "fmt"

func main() {

	var firstName, lastName string = "muhammad", "ivan"

	var myName, myBro = "ivan", "john"

	x := "Hallo"
	y := 25

	angka := 15.5
	txt := "hello world!"

	fmt.Print(firstName, " ", lastName, "\n")
	fmt.Println(myName, myBro)

	fmt.Printf("value x adalah : %v dan tipe data x adalah : %T\n", x, x)
	fmt.Printf("value y adalah : %v dan tipe data y adalah : %T\n", y, y)

	fmt.Printf("%v%%\n", angka)
	fmt.Printf("%#v", txt)
}
