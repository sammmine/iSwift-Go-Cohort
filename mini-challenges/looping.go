//mini challenge 1 - fizzbuzz
//looping and conditional

package looping

import "fmt"

func looping() {
	n := 15
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			if i%5 == 0 {
				fmt.Println("Buzz")
			} else if i%3 == 0 {
				fmt.Println("Fizz")
			} else {
				fmt.Println(i)
			}
		}
	}
}
