package main

import "fmt"

func main() {

	var myBrother string
	myBrother = "Poter"
	var myName string = "John doe"
	var nickName = "John"
	myAge := 28

	// default variabel
	var a string
	var b int
	var c bool

	// deklarasi multiple variabel
	var d, e, f, g = 1, 2, 3, "Hello World!"

	// variabel dalam blok
	var (
		firstName string = "ivan"
		lastName         = "satria"
		height    int    = 176
	)

	fmt.Println(myName)
	fmt.Println(nickName)
	fmt.Println(myBrother)
	fmt.Println(myAge)

	// default variabel
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// deklarasi multiple variabel
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// var dalam blok
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(height)

}
