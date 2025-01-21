package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"ivan", "satria"}
	printMessage("hallo", names)
}

func printMessage(message string, arr []string) {
	namesString := strings.Join(arr, " ")
	fmt.Println(message, namesString)
}
