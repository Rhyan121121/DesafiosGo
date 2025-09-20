package main

import "fmt"

func DivisibleBy3(x int) bool {
	if x%3 == 0 {
		return true
	}

	return false

}

func DivisibleBy5(x int) bool {
	if x%5 == 0 {
		return true
	}
	return false
}

func main() {
	for i := 1; i <= 100; i++ {
		if DivisibleBy3(i) {
			fmt.Println("Pin")
			continue
		} else if DivisibleBy5(i) {
			fmt.Println("Pan")
			continue
		}
		fmt.Println(i)
	}
}
