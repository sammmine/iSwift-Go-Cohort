// Mini Challenge 2 : Calculating Words Appear

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// input menggunakan bufio package
	fmt.Println("Masukkan kata: ")
	scn := bufio.NewReader(os.Stdin)
	kata, _ := scn.ReadString('\n')

	// membuat map untuk menampung jumlah kemunculan huruf
	// abjad sebagai key, dan jumlah kemunculan sebagai value
	jumlahHuruf := make(map[rune]int)

	// memasukan key dan value ke dalam map
	for _, huruf := range kata {
		jumlahHuruf[huruf]++
	}

	// output dengan looping map
	fmt.Printf("\nJumlah kemunculan huruf:\n")
	for huruf, jumlah := range jumlahHuruf {
		fmt.Printf("%c : %d \n", huruf, jumlah)
	}
}
