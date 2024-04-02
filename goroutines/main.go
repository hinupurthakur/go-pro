package goroutines

// You can edit this code!
// Click here and start typing.

import (
	"fmt"
	"time"
)

// Problem demonstrates the problem with goroutines when using a loop variable
func Problem() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range values {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(5 * time.Second)
}

func Solution() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range values {
		go func(value int) {
			fmt.Println(value)
		}(v)
	}
	time.Sleep(5 * time.Second)
}
