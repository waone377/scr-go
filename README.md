# Pengikis Web Go (Go Web Scraper)

Sebuah alat pengikis web (web scraper) interaktif berbasis baris perintah (command-line) yang dibuat dengan Go. Alat ini dirancang agar aman, efisien, dan ramah pengguna, memungkinkan Anda untuk mengekstrak data dari situs web dengan mudah sambil mencoba menghindari deteksi bot.

## Fitur

- **Antarmuka Berbasis Teks**: Pengguna dipandu melalui serangkaian prompt yang jelas untuk memasukkan URL target, memilih elemen HTML, dan memilih format output menggunakan input numerik.
- **Aman & Tersembunyi**: Secara otomatis merotasi header User-Agent dan menerapkan penundaan acak di antara permintaan untuk meniru perilaku penjelajahan manusia dan mengurangi kemungkinan diblokir.
- **Pemilihan Elemen Fleksibel**: Pengguna dapat memilih beberapa elemen HTML (seperti `h1`, `p`, `a`, dll.) untuk diekstrak dari halaman target.
- **Opsi Output Ganda**: Pilih untuk mengekstrak hanya teks bersih dari elemen atau mendapatkan konten HTML penuh, lengkap dengan tag dan atribut.
- **Output Terstruktur**: Hasil scraping disimpan dalam file JSON yang terformat dengan baik, dengan nama file unik berdasarkan timestamp, membuatnya mudah untuk di-parse dan digunakan lebih lanjut.
- **Struktur Proyek Modular**: Kode diorganisir ke dalam paket-paket yang berbeda (`antarmuka`, `pengikis`, `utilitas`), membuatnya mudah untuk dipelihara dan diperluas.

## Prasyarat

- Go versi 1.21 atau lebih tinggi.

## Instalasi

1.  **Clone repository:**
    ```sh
    git clone <url-repository-anda> pengikis-web-go
    cd pengikis-web-go
    ```

2.  **Install dependensi:**
    Alat ini menggunakan Go Modules. Dependensi akan diunduh secara otomatis saat Anda membangun atau menjalankan proyek.
    ```sh
    go mod tidy
    ```

## Cara Penggunaan

Jalankan aplikasi dari direktori root proyek:

```sh
go run ./cmd/pengikisweb
```

Alat ini akan memulai sesi interaktif:

1.  **Masukkan URL Target**: Anda akan diminta untuk memasukkan URL lengkap dari halaman web yang ingin Anda kikis.
    ```
    Silakan masukkan URL target:
    > <ketik-url-anda-di-sini>
    ```

2.  **Pilih Elemen HTML**: Sebuah daftar elemen HTML bernomor akan muncul. Masukkan nomor yang sesuai, dipisahkan dengan koma.
    ```
    Pilih elemen HTML yang akan dikikis (pisahkan dengan koma, contoh: 1, 3, 8):
    1. h1
    2. h2
    3. p
    ...
    > 1, 3, 8
    ```

3.  **Pilih Format Output**: Tentukan apakah Anda ingin mengekstrak hanya teks bersih atau tag HTML penuh dengan memasukkan nomor pilihan.
    ```
    Pilih format keluaran:
    1. Teks Bersih Saja
    2. HTML Penuh
    > 1
    ```

Setelah Anda menyelesaikan semua prompt, scraper akan memulai prosesnya. Setelah selesai, pesan sukses akan ditampilkan, menunjukkan nama file JSON tempat data Anda disimpan.

## Struktur Output

Hasilnya akan disimpan dalam file bernama `hasil_<timestamp>.json` di direktori root proyek. Strukturnya adalah sebuah objek JSON di mana setiap kunci sesuai dengan pemilih HTML yang Anda pilih, dan nilainya adalah array string yang berisi data yang diekstraksi.

**Contoh `hasil_1679812345.json`:**
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

## Penafian

Harap gunakan alat ini secara bertanggung jawab. Hormati file `robots.txt` situs web dan syarat layanan mereka. Scraping yang agresif dapat membebani server situs web dan dapat menyebabkan alamat IP Anda diblokir. Pengembang tidak bertanggung jawab atas penyalahgunaan alat ini.