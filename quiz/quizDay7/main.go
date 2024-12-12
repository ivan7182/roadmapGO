package main

import "fmt"

/*
Di sebuah perpustakaan, terdapat sebuah sistem pencatatan daftar buku yang sedang dipinjam.
Sistem ini mencatat nama-nama buku dalam bentuk slice string. Namun,
ada beberapa aturan yang harus diterapkan:

1. Jika seorang pengguna mengembalikan buku, maka nama buku tersebut harus dihapus dari daftar.
2. Jika seorang pengguna meminjam buku baru, nama buku tersebut harus ditambahkan ke daftar.
3. Perpustakaan ingin mencatat jumlah buku yang saat ini sedang dipinjam.

*/

func main() {
	soal1()
}

func soal1() {
	book := []string{"Go Programming", "Database Design", "Machine Learning"}
	fmt.Println("daftar buku :", book)

	//penambahan buku
	addBook := append(book, "Web Development")
	fmt.Println("Daftar buku setelah penambahan:", addBook)

	//pengghapusan buku
	addBook = removeBook(addBook, "Database Design")
	fmt.Println("Daftar buku setelah penghapusan:", addBook)

	//
	fmt.Println(len(addBook))

}

func removeBook(slice []string, book string) []string {
	for i, b := range slice {
		if b == book {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
