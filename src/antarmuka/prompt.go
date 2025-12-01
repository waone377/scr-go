package antarmuka

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

func BersihkanLayar() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}

func DapatkanURLTarget() string {
	pembaca := bufio.NewReader(os.Stdin)
	var url string
	for {
		fmt.Println("Silakan masukkan URL target:")
		fmt.Print("> ")
		url, _ = pembaca.ReadString('\n')
		url = strings.TrimSpace(url)
		if url != "" {
			break
		}
		fmt.Println("URL tidak boleh kosong. Silakan coba lagi.")
	}
	return url
}

func DapatkanElemenTarget() ([]string, error) {
	opsiElemen := []string{"h1", "h2", "h3", "p", "a", "span", "div", "li", "ul", "ol", "img", "table", "tr", "td", "th", "form", "input", "button", "body", "html"}
	pembaca := bufio.NewReader(os.Stdin)

	fmt.Println("Pilih elemen HTML yang akan dikikis (pisahkan dengan koma, contoh: 1, 3, 8):")
	for i, elemen := range opsiElemen {
		fmt.Printf("%d. %s\n", i+1, elemen)
	}

	var elemenTerpilih []string
	for {
		fmt.Print("> ")
		input, _ := pembaca.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Input tidak boleh kosong. Silakan coba lagi.")
			continue
		}

		potongan := strings.Split(input, ",")
		indeksTerpilih := make(map[int]bool)
		elemenTerpilih = []string{}
		valid := true

		for _, bagian := range potongan {
			nomor, err := strconv.Atoi(strings.TrimSpace(bagian))
			if err != nil || nomor < 1 || nomor > len(opsiElemen) {
				fmt.Printf("Input tidak valid: '%s'. Harap masukkan nomor antara 1 dan %d.\n", bagian, len(opsiElemen))
				valid = false
				break
			}
			indeks := nomor - 1
			if !indeksTerpilih[indeks] {
				elemenTerpilih = append(elemenTerpilih, opsiElemen[indeks])
				indeksTerpilih[indeks] = true
			}
		}

		if valid && len(elemenTerpilih) > 0 {
			break
		}
	}
	return elemenTerpilih, nil
}

func DapatkanFormatKeluaran() (string, error) {
	opsiFormat := []string{"Teks Bersih Saja", "HTML Penuh"}
	pembaca := bufio.NewReader(os.Stdin)
	var format string

	fmt.Println("Pilih format keluaran:")
	for i, opsi := range opsiFormat {
		fmt.Printf("%d. %s\n", i+1, opsi)
	}

	for {
		fmt.Print("> ")
		input, _ := pembaca.ReadString('\n')
		input = strings.TrimSpace(input)
		nomor, err := strconv.Atoi(input)

		if err != nil || nomor < 1 || nomor > len(opsiFormat) {
			fmt.Printf("Input tidak valid. Harap masukkan nomor antara 1 dan %d.\n", len(opsiFormat))
			continue
		}
		format = opsiFormat[nomor-1]
		break
	}
	return format, nil
}
