package main

import (
	"fmt"
	"sync"
)

// Untuk yang arranged, saya tidak mengerti jika menggunakan tipe data interface.
// Sehingga di sini untuk kedua datanya saya memakai tipe data struct.
// Isi dari struct adalah slice of string (untuk datanya) dan sync.Mutex.
type data struct {
	data  []string
	mutex sync.Mutex
}

// Method ini untuk menampilkan/mengoutput isi struct.
// Mutex diimplementasikan agar output terurut (atau synchronize).
func (d *data) getData() string {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return fmt.Sprintf("[%s %s %s]", d.data[0], d.data[1], d.data[2])
}

func main() {
	// Pertama kita isi kedua buah data yang bertipe struct.
	slice1 := &data{data: []string{"coba1", "coba2", "coba3"}}
	slice2 := &data{data: []string{"bisa1", "bisa2", "bisa3"}}

	var wg sync.WaitGroup

	wg.Add(1)

	// Implementasi concurrency di sini saya hanya menggunakan satu goroutine.
	go func() {
		defer wg.Done()
		for i := 1; i <= 4; i++ {
			fmt.Println(slice1.getData(), i)
			fmt.Println(slice2.getData(), i)
		}
	}()

	wg.Wait()
}
