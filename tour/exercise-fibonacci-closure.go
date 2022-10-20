package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	oneBehind := 1
	twoBehind := 0
	n := 0
	return func() int {
		if n == 0 {
			n++
			return 0
		} else if n == 1 {
			n++
			return 1
		}
		ret := oneBehind + twoBehind
		twoBehind = oneBehind
		oneBehind = ret
		n++
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
