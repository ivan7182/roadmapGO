package main

import "fmt"

func main() {
	var variableName1 string = "hello wordl"
	variableName2 := "hayyy"

	fmt.Println(variableName1)
	fmt.Println(variableName2)

	//primitive: boolean, int, float, string

	//boolean
	boolVar := true
	fmt.Println("type : %T value: %v\n", boolVar, boolVar)

	//integer
	intVar := int(4)

	fmt.Println("type: %T value: %v\n", intVar, intVar)

	//float
	floatVar1 := float32(3.5)
	floatVar2 := float32(4.8)
	fmt.Println("type: %T value: %v\n", floatVar1, floatVar1)
	fmt.Println("type: %t value: %v\n", floatVar2, floatVar2)

	//bytes
	bytesVar := byte(255)
	fmt.Println("type: %T value: %v\n", bytesVar, bytesVar)

	bytesVar2 := []byte("hello wordl")
	fmt.Println("type: %T value: %v\n", bytesVar2, bytesVar2)

	//complex
	complexVar := -7 + 3i
	fmt.Println("type %T value: %v\n", complexVar, complexVar)

	//interface
	var myInterfaceVar interface{}

	myInterfaceVar = 8
	fmt.Printf("type: %T value: %v\n", myInterfaceVar, myInterfaceVar)
	myInterfaceVar = "Hello wordl"
	fmt.Printf("type: %T value: %v\n", myInterfaceVar, myInterfaceVar)
}

// type MethodList interface{
// 	myFunction()
// 	myFunction2(int) int
// }
