# Analisis Output:

Perhatikan output yang dihasilkan dan bandingkan dengan informasi di dokumen PDF.
Pastikan Anda memahami setiap bagian dari output tersebut.

## Source
    func main() {
    // Mendapatkan waktu sekarang
        now := time.Now()
        fmt.Println("Waktu sekarang:", now)

    // Menggunakan berbagai method time.Time
        fmt.Println("Tahun:", now.Year())
        fmt.Println("Bulan:", now.Month())
        fmt.Println("Tanggal:", now.Day())
        fmt.Println("Jam:", now.Hour())
        fmt.Println("Menit:", now.Minute())
        fmt.Println("Detik:", now.Second())

    // Parsing waktu dari string
        layoutFormat := "2006-01-02 15:04:05"
        value := "2015-09-02 08:04:00"
        date, _ := time.Parse(layoutFormat, value)
        fmt.Println("Waktu yang di-parse:", date)
    }

## go run main.go
    Waktu sekarang: 2024-11-08 10:27:12.4110584 +0700 +07 m=+0.000000001
    Tahun: 2024
    Bulan: November
    Tanggal: 8
    Jam: 10
    Menit: 27
    Detik: 12
    Waktu yang di-parse: 2015-09-02 08:04:00 +0000 UTC

# Eksperimen:

Coba modifikasi kode untuk menggunakan fungsi-fungsi time lainnya yang disebutkan dalam dokumen.
Misalnya, tambahkan penggunaan time.Date() untuk membuat waktu tertentu.



# Refleksi:

Tuliskan catatan singkat tentang apa yang Anda pelajari dan bagian mana yang masih perlu dipelajari lebih lanjut.