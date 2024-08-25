// Mini Challenge 3 : Struct, Slice, Map, Function

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Peserta struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func Ketemu(target Peserta) {
	if target.Nama != "" {
		fmt.Println("Absen:", target.Absen)
		fmt.Println("Nama:", target.Nama)
		fmt.Println("Alamat:", target.Alamat)
		fmt.Println("Pekerjaan:", target.Pekerjaan)
		fmt.Println("Alasan:", target.Alasan)
	} else {
		fmt.Printf("Data tidak ditemukan.")
	}
}

func main() {
	// DATA
	var pesertaAll = []Peserta{
		{1, "HaloHalo", "Jl. Lorem", "Backend", "Alasan"},
		{2, "HaiHai", "Jl. Ipsum", "UI/UX", "Alasan"},
	}

	if len(os.Args) < 2 {
		// KONDISI TIDAK ADA ARGUMEN

		fmt.Println("Tolong masukan nama atau nomor absen!")
		fmt.Println("Contoh: 'go run main.go HaloHalo' atau 'go run main.go 1'")

	} else {
		// KONDISI ADA ARGUMEN

		argumen := os.Args[1]

		// MEMBEDAKAN ARGUMEN BERTIPE INTEGER (ABSEN) DAN STRING (NAMA)
		hasil, err := strconv.Atoi(argumen)
		if err != nil {
			// ARGUMEN BERTIPE STRING (FIELD NAMA)

			var target Peserta

			for _, peserta := range pesertaAll {
				if peserta.Nama == argumen {
					target = peserta
					break
				}
			}

			Ketemu(target) // JIKA PESERTA TIDAK DITEMUKAN, VARIABEL pesertaDitemukan KOSONG ("")

		} else {
			// ARGUMEN BERTIPE INTEGER (FIELD ABSEN)

			var target Peserta

			for _, peserta := range pesertaAll {
				if peserta.Absen == hasil {
					target = peserta
					break
				}
			}

			Ketemu(target) // JIKA PESERTA TIDAK DITEMUKAN, VARIABEL pesertaDitemukan KOSONG ("")

		}
	}
}
