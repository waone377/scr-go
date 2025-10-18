# scr-go — Pengikis Web Interaktif (Go)

scr-go adalah alat pengikis web (web scraper) berbasis baris perintah yang ditulis dengan Go. Aplikasi ini bersifat interaktif: pengguna dipandu melalui serangkaian prompt untuk memasukkan URL target, memilih elemen HTML yang ingin diekstrak, serta menentukan format keluaran. Tujuan utamanya adalah membuat scraping yang aman, dapat diulang, dan mudah dipakai.

## Fitur utama

- Antarmuka teks interaktif yang memandu pengguna langkah-demi-langkah.
- Rotasi header User-Agent otomatis dan penundaan acak antar permintaan untuk meniru perilaku penjelajahan manusia.
- Pilihan elemen HTML fleksibel (mis. h1, h2, p, a, dll.) — bisa memilih beberapa elemen sekaligus.
- Pilihan format keluaran: teks bersih (tanpa tag HTML) atau HTML penuh (dengan tag dan atribut).
- Hasil disimpan sebagai file JSON terformat dengan nama unik berdasarkan timestamp.
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
go run ./cmd/pengikisweb
```

- Membangun binary lalu menjalankan:
```sh
go build -o scr-go ./cmd/pengikisweb
./scr-go
```

Catatan: Nama package / entrypoint di repo ini adalah `cmd/pengikisweb`. Jika Anda mengganti struktur, sesuaikan perintah di atas.

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

Hasil scraping disimpan di file `hasil_<timestamp>.json` di direktori root proyek (timestamp dalam format Unix epoch). Format JSON:

- Kunci: nama elemen/pemilih (selector) yang dipilih.
- Nilai: array string berisi hasil ekstraksi untuk selector tersebut.
- Jika memilih "Teks Bersih", setiap item adalah teks (innerText) dari elemen.
- Jika memilih "HTML Penuh", setiap item adalah HTML dari elemen (outerHTML).

Contoh `hasil_1679812345.json`:
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

## Konfigurasi (opsional)

Catatan: Implementasi detil konfigurasi tergantung pada kode. Beberapa opsi yang biasanya tersedia atau mudah ditambahkan:
- Menentukan daftar User-Agent kustom.
- Mengatur rentang delay acak (mis. 1–5 detik).
- Mengatur timeout HTTP dan retry policy.
- Menyaring hasil berdasarkan atribut (mis. hanya <a> dengan atribut target tertentu).

Jika Anda ingin saya menambahkan file konfigurasi (mis. TOML/JSON) dan parsingnya di kode, beri tahu saya dan saya bisa membuat PR contoh.

## Struktur proyek (sekilas)

- cmd/pengikisweb — entrypoint aplikasi.
- internal/ atau pkg/ (mis. antarmuka, pengikis, utilitas) — modul fungsional (nama folder bisa berbeda tergantung implementasi).
- go.mod — manajemen dependensi.
- hasil_*.json — contoh file keluaran yang dihasilkan oleh program.

Sesuaikan struktur direktori jika repo Anda berbeda; README ini dimaksudkan untuk memberikan gambaran umum yang jelas.

## Kontribusi

Kontribusi welcome! Jika Anda ingin:
- Menambahkan fitur (mis. support untuk paging, autentikasi, proxy),
- Memperbaiki bug,
- Menambahkan pengujian unit atau integrasi,

Buka issue atau kirim pull request (PR) dengan deskripsi perubahan.

## Lisensi

Tambahkan file LICENSE di repo dengan lisensi yang Anda pilih (mis. MIT, Apache-2.0). README ini tidak menyertakan lisensi secara otomatis.

## Penafian

Gunakan alat ini secara bertanggung jawab. Penulis/pemelihara tidak bertanggung jawab atas penggunaan yang melanggar hukum atau yang menimbulkan kerusakan pada pihak ketiga.

---
Terima kasih telah menggunakan scr-go — jika Anda ingin, saya bisa:
- Menyelaraskan README ini secara langsung dengan struktur file di repo (mencantumkan daftar file/struktur sebenarnya),
- Menambahkan contoh config yang dapat dibaca program, atau
- Membuat file CONTRIBUTING.md dan contoh issue/PR template.
Pilih salah satu dan saya akan menyiapkannya. 
