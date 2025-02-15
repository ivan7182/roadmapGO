package main

import "fmt"

func incrementPtr(var1 *int) {
	*var1 = *var1 + 1
}
func increment(var1 int) {
	var1 = var1 + 1
}
func main() {
	var1 := 1
	fmt.Printf("my value is %d\n", var1)
	increment(var1)
	fmt.Printf("my value is %d\n", var1)
	incrementPtr(&var1)
	fmt.Printf("my value is %X, value is %d\n", &var1, var1)

	var var2 *int
	fmt.Printf("my value is %d\n", var2)

	var2 = new(int)
	*var2 = 5
	fmt.Printf("my value is %d\n", *var2)

	var2 = &var1
	*var2 = 899
	fmt.Printf("my value is %d\n", var1)
}
