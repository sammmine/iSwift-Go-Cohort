// Mini Challenge 4 : Interface & Goroutine

package main

import (
	"fmt"
	"sync"
)

type dataa interface {
}

func rutinn(slice dataa, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(slice, i)
}

func main() {
	var wg sync.WaitGroup

	var slice1 dataa = []dataa{"coba1", "coba2", "coba3"}
	var slice2 dataa = []dataa{"bisa1", "bisa2", "bisa3"}

	wg.Add(8)

	for i := 1; i < 5; i++ {
		go rutinn(slice1, i, &wg)
		go rutinn(slice2, i, &wg)
	}

	wg.Wait()

}
