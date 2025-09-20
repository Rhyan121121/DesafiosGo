package main

import "fmt"

func Divisible(x int) bool {
	if x%3 == 0 {
		return true
	}

	return false

}

func main() {
	for i := 1; i <= 100; i++ {
		if Divisible(i) {
			fmt.Println(i)
		}
	}
}
