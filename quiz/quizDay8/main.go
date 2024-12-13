package main

import "fmt"

/*
1 :
Buat sebuah map yang menyimpan nama negara sebagai kunci dan ibu kotanya sebagai nilai. Kemudian,
buat fungsi yang menerima nama
negara sebagai input dan mencetak ibu kotanya. Jika negara tidak ditemukan, cetak pesan: "Negara tidak ditemukan."

2:
Diberikan sebuah map dengan kunci berupa nama siswa dan nilai berupa skor ujiannya. Tugas Anda adalah:
- Menambahkan seorang siswa baru dengan skor ujian tertentu.
- Menghapus data seorang siswa berdasarkan namanya.
- Mencari skor siswa berdasarkan nama.
- Menampilkan semua siswa dan skor mereka.

3:

Buat sebuah program untuk menyimpan data jadwal mata pelajaran. Map memiliki struktur:
go
map[string][]string{"hari": {"mata pelajaran"}}

Tugas:
- Tambahkan mata pelajaran baru pada hari tertentu.
- Hapus mata pelajaran dari hari tertentu.
- Tampilkan jadwal seluruh hari dan mata pelajarannya.

*/
func main() {
	//soal1()
	//soal2()
	soal3()

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

func soal2() {
	var sekolah = map[string]int{
		"ivan":   90,
		"satria": 86,
	}

	fmt.Println("data siswa:", sekolah)

	//tambah siswa
	sekolah["john"] = 90
	sekolah["doe"] = 85

	fmt.Println("data siswa setelah ditambah:", sekolah)

	// hapus siswa
	delete(sekolah, "john")
	fmt.Println("hapus siswa:", sekolah)

	//cari siswa
	cariSiswa := "bajang"
	value, isExist := sekolah[cariSiswa]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("siswa belum terdaftar")
	}

}

func soal3() {
	mapel := []map[string]string{
		{"hari": "senin", "mata pelajaran": "matematika"},
		{"hari": "selasa", "mata pelajaran": "ipa"},
		{"hari": "rabu", "mata pelajaran": "ips"},
	}
	fmt.Println("jadwal pelajaran :", mapel)

	jadwalBaru := map[string]string{
		"hari": "kamis", "mata pelajaran": "ipa",
	}
	mapel = append(mapel, jadwalBaru)

	fmt.Println("jadwal baru ditambahkan:", mapel)

	mapel = deleteMapel(mapel, "selasa")

	fmt.Println("jadwal settelah dihapus :", mapel)
}

func deleteMapel(mapel []map[string]string, hari string) []map[string]string {
	for i, m := range mapel {
		if m["hari"] == hari {
			return append(mapel[:i], mapel[i+1:]...)
		}
	}
	fmt.Println("hari tidak ditemukan:", hari)
	return mapel
}
