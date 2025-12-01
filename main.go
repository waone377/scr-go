package main

import (
	"fmt"
	"log"
	"time"

	"pengikisweb/src/antarmuka"
	"pengikisweb/src/pengikis"
	"pengikisweb/src/utilitas"
)

func main() {
	urlTarget := antarmuka.DapatkanURLTarget()
	antarmuka.BersihkanLayar()

	elemenTarget, err := antarmuka.DapatkanElemenTarget()
	if err != nil {
		log.Fatalf("Kesalahan mendapatkan elemen target: %v", err)
	}
	antarmuka.BersihkanLayar()

	formatKeluaran, err := antarmuka.DapatkanFormatKeluaran()
	if err != nil {
		log.Fatalf("Kesalahan mendapatkan format keluaran: %v", err)
	}
	antarmuka.BersihkanLayar()

	fmt.Printf("Mengikis %s...\n", urlTarget)

	data, err := pengikis.Kikis(urlTarget, elemenTarget, formatKeluaran)
	if err != nil {
		log.Fatalf("Pengikisan gagal: %v", err)
	}

	if len(data) == 0 {
		fmt.Println("Tidak ada data yang terkikis. Elemen target mungkin tidak ada di halaman tersebut.")
		return
	}

	namaBerkasKeluaran := fmt.Sprintf("output/hasil_%d.json", time.Now().Unix())
	err = utilitas.SimpanKeJSON(data, namaBerkasKeluaran)
	if err != nil {
		log.Fatalf("Gagal menyimpan data ke JSON: %v", err)
	}

	fmt.Printf("Pengikisan berhasil! Data disimpan ke %s\n", namaBerkasKeluaran)
}
