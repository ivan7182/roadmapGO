package main

import "fmt"

/*
1 :
Buat sebuah map yang menyimpan nama negara sebagai kunci dan ibu kotanya sebagai nilai. Kemudian,
buat fungsi yang menerima nama
negara sebagai input dan mencetak ibu kotanya. Jika negara tidak ditemukan, cetak pesan: "Negara tidak ditemukan."


*/
func main() {
	soal1()

}

func soal1() {
	negara := map[string]string{
		"Indonesia": "jakarta",
		"amerika":   "Washington, D.C.",
		"brasil":    "Brasilia",
	}

	cariNegara := "amerika"
	value, isExist := negara[cariNegara]

	if isExist {
		fmt.Printf("ibu kota %s adalah %s. \n ", cariNegara, value)
	} else {
		fmt.Println("Negara tidak ditemukan.")
	}
}
