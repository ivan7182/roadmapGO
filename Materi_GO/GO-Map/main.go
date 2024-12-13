package main

import "fmt"

func main() {
	// gomap()
	// inisialisasiMap()
	// iterasiMap()
	// keyMap()
	sliceAndMap()
}
func gomap() {
	sembako := map[string]string{}
	sembako["beras"] = "1kg"
	sembako["gula"] = "2kg"

	fmt.Println("ketersedian sembako :", sembako["beras"])
	fmt.Println("garam ", sembako["gaaram"])
}

func inisialisasiMap() {
	sembako := map[string]int{
		"beras": 1,
		"gula":  3,
	}
	fmt.Println(sembako)
}

func iterasiMap() {
	sembako := map[string]string{
		"beras":  "1kg",
		"gula":   "2kg",
		"terigu": "4kg",
	}

	for key, val := range sembako {
		fmt.Println(key, val)
	}

	delete(sembako, "gula")
	fmt.Println(len(sembako))
	fmt.Println(sembako)
}

func keyMap() {
	sembako := map[string]string{
		"beras":   "1kg",
		"gula":    "3kg",
		"teriigu": "5kg",
	}

	val, isExist := sembako["gula"]

	if isExist {
		fmt.Println(val)
	} else {
		fmt.Println("eror")
	}
}

func sliceAndMap() {
	sembako := []map[string]string{
		map[string]string{"nama": "terigu", "jenis": "pangan"},
		map[string]string{"nama": "beras", "jenis": "pangan"},
		map[string]string{"nama": "gaaram", "jenis": "panngan"},
	}

	for _, sembako := range sembako {
		fmt.Println(sembako["jenis"], sembako["nama"])
	}
}
