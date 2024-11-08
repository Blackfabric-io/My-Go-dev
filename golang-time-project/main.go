package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Waktu sekarang:", now)

	fmt.Println("Tahun:", now.Year())
	fmt.Println("Bulan:", now.Month())
	fmt.Println("Tanggal:", now.Day())
	fmt.Println("Jam:", now.Hour())
	fmt.Println("Menit:", now.Minute())
	fmt.Println("Detik:", now.Second())

	layoutFormat := "2006-01-02 15:04:05"
	value := "2015-09-02 08:04:00"
	date, _ := time.Parse(layoutFormat, value)
	fmt.Println("Waktu yang di-parse:", date)
}
