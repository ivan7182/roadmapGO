package main

import "fmt"

func main() {
	// slice()
	fungsi()
}

func slice() {
	sembako := []string{"gula", "beras", "terigu", "garam"}
	// newSembako := sembako[0:]
	fmt.Println(len(sembako))

	// var aSembako = sembako[0:3]
	// var bSembako = sembako[1:4]

	// var aasembako = aSembako[1:2]
	// var bbsembako = bSembako[0:1]

	// fmt.Println(aSembako)
	// fmt.Println(bSembako)
	// fmt.Println(len(aasembako))
	// fmt.Println(bbsembako)

	// bbsembako[0] = "beras baru"

	// fmt.Println(aSembako)
	// fmt.Println(bSembako)
	// fmt.Println(aasembako)
	// fmt.Println(bbsembako)
}

func fungsi() {
	sembako := []string{"gula", "beras", "terigu"}
	// fmt.Println(len(sembako))
	// fmt.Println(cap(sembako))

	aSembako := sembako[0:2]
	fmt.Println(len(aSembako))
	fmt.Println(cap(aSembako))

	addSembako := append(sembako, "garam")
	fmt.Println("setelah ditambah :", addSembako)

}
