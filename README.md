# scr-go — Pengikis Web Interaktif (Go)

scr-go adalah alat pengikis web (web scraper) berbasis baris perintah yang ditulis dengan Go. Aplikasi ini bersifat interaktif: pengguna dipandu melalui serangkaian prompt untuk memasukkan URL target, memilih elemen HTML yang ingin diekstrak, serta menentukan format keluaran. Tujuan utamanya adalah membuat scraping yang aman, dapat diulang, dan mudah dipakai.

## Fitur utama

- Antarmuka teks interaktif yang memandu pengguna langkah-demi-langkah.
- Rotasi header User-Agent otomatis dan penundaan acak antar permintaan untuk meniru perilaku penjelajahan manusia.
- Pilihan elemen HTML fleksibel (mis. h1, h2, p, a, dll.) — bisa memilih beberapa elemen sekaligus.
- Pilihan format keluaran: teks bersih (tanpa tag HTML) atau HTML penuh (dengan tag dan atribut).
- Hasil disimpan sebagai file JSON terformat di dalam folder `output` dengan nama unik berdasarkan timestamp.
- Struktur kode modular untuk memudahkan pemeliharaan dan pengembangan.

## Prasyarat

- Go 1.21 atau lebih tinggi terinstal pada sistem Anda.
- Koneksi internet untuk mengunduh dependensi dan melakukan permintaan HTTP.

## Instalasi & Persiapan

1. Clone repository:
```sh
git clone https://github.com/waone377/scr-go.git
cd scr-go
```

2. Unduh dependensi (Go Modules):
```sh
go mod tidy
```

## Menjalankan (Run)

- Menjalankan langsung dari source:
```sh
go run main.go
```

- Membangun binary lalu menjalankan:
```sh
go build -o scr-go main.go
./scr-go
```

## Cara penggunaan (interaktif)

Setelah menjalankan program, Anda akan diminta memasukkan beberapa input:

1. Masukkan URL target:
```
Silakan masukkan URL target:
> https://example.com
```

2. Pilih elemen HTML yang ingin diekstrak. Program akan menampilkan daftar bernomor, misal:
```
Pilih elemen HTML yang akan dikikis (pisahkan dengan koma, contoh: 1,3,8):
1. h1
2. h2
3. p
4. a
5. img
> 1,3,4
```

3. Pilih format output:
```
Pilih format keluaran:
1. Teks Bersih Saja (hanya inner text)
2. HTML Penuh (termasuk tag dan atribut)
> 1
```

Setelah input selesai, scraper akan mulai memproses halaman sesuai pilihan. Selama proses, program menerapkan rotasi User-Agent dan jeda acak antara permintaan untuk mengurangi risiko pemblokiran.

## Struktur file keluaran

Hasil scraping disimpan di file `output/hasil_<timestamp>.json` di direktori root proyek (timestamp dalam format Unix epoch). Format JSON:

- Kunci: nama elemen/pemilih (selector) yang dipilih.
- Nilai: array string berisi hasil ekstraksi untuk selector tersebut.
- Jika memilih "Teks Bersih", setiap item adalah teks (innerText) dari elemen.
- Jika memilih "HTML Penuh", setiap item adalah HTML dari elemen (outerHTML).

Contoh `output/hasil_1679812345.json`:
```json
{
  "h1": [
    "Judul Utama Halaman"
  ],
  "p": [
    "Ini adalah paragraf pertama.",
    "Ini adalah paragraf kedua."
  ],
  "a": [
    "<a href=\"/link1\">Teks Link 1</a>",
    "<a href=\"/link2\" class=\"external\">Teks Link 2</a>"
  ]
}
```

## Perilaku Permintaan & Etika

- scr-go mencoba meniru perilaku manusia dengan merotasi header User-Agent dan menambahkan jeda acak. Ini bukanlah jaminan untuk menghindari pemblokiran.
- Hormati file `robots.txt` dan ketentuan layanan situs web yang Anda kunjungi. Jangan melakukan scraping agresif yang dapat membebani server.
- Scraping konten berhak cipta atau pribadi tanpa izin dapat melanggar hukum — pastikan penggunaan Anda sesuai hukum yang berlaku.

## Struktur proyek

- `main.go` — Titik masuk (entrypoint) utama aplikasi.
- `src/` — Berisi paket-paket logika inti aplikasi.
  - `antarmuka/` — Logika untuk interaksi dengan pengguna via CLI.
  - `pengikis/` — Logika inti untuk melakukan scraping web.
  - `utilitas/` — Fungsi bantuan, seperti menyimpan file.
- `output/` — Direktori tempat hasil scraping disimpan.
- `go.mod`, `go.sum` — Manajemen dependensi Go.

## Kontribusi

Kontribusi welcome! Jika Anda ingin:
- Menambahkan fitur (mis. support untuk paging, autentikasi, proxy),
- Memperbaiki bug,
- Menambahkan pengujian unit atau integrasi,

Buka issue atau kirim pull request (PR) dengan deskripsi perubahan.

## Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT. Lihat file `LICENSE` untuk detailnya.

## Penafian

Gunakan alat ini secara bertanggung jawab. Penulis/pemelihara tidak bertanggung jawab atas penggunaan yang melanggar hukum atau yang menimbulkan kerusakan pada pihak ketiga.