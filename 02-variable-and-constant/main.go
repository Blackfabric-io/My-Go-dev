package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel dengan tipe data
	var text1 string
	text1 = "ini teks 1"
	fmt.Println(text1)

	var text2 string = "ini teks 2"
	text2 = "ini teks 1 diubah"
	fmt.Println(text2)

	// Deklarasi variabel tanpa tipe data menggunakan :=
	text3 := "ini teks 3"
	text3 = "ini teks 3 diubah"
	fmt.Println(text3)

	text4 := "ini teks 4"
	fmt.Println(text4)

	// Deklarasi konstanta
	const judul = "Ini Judul"
	fmt.Println(judul)

	// Jika Anda mencoba mengubah nilai konstanta, akan terjadi error
	// judul = "judul di ubah" // Uncommenting this line will cause an error
}
