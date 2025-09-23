package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func waitAndPrint(c chan string) {
	for {
		message := <-c
		fmt.Println(message)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	var c chan string = make(chan string)
	go ping(c)
	go pong(c)
	go waitAndPrint(c)

	var stop string
	fmt.Scanln(&stop)

}
