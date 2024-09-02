package main

import (
	"fmt"
	"sync"
)

type dataa interface {
}

func rutinn(slice dataa, i int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	fmt.Println(slice, i)
	mutex.Unlock()
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	var slice1 dataa = []dataa{"coba1", "coba2", "coba3"}
	var slice2 dataa = []dataa{"bisa1", "bisa2", "bisa3"}

	wg.Add(8)

	for i := 1; i < 5; i++ {
		mutex.Lock()
		go rutinn(slice1, i, &wg, &mutex)
		go rutinn(slice2, i, &wg, &mutex)
		mutex.Unlock()
	}

	wg.Wait()

}
